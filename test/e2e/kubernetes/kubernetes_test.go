// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/engine"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/deployment"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/hpa"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/job"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/namespace"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/networkpolicy"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/node"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/persistentvolume"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/persistentvolumeclaims"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/pod"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/service"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/storageclass"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/Azure/aks-engine/test/e2e/remote"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	WorkloadDir                               = "workloads"
	PolicyDir                                 = "workloads/policies"
	retryCommandsTimeout                      = 5 * time.Minute
	kubeSystemPodsReadinessChecks             = 6
	sleepBetweenRetriesWhenWaitingForPodReady = 1 * time.Second
	sleepBetweenRetriesRemoteSSHCommand       = 3 * time.Second
	timeoutWhenWaitingForPodOutboundAccess    = 1 * time.Minute
	stabilityCommandTimeout                   = 3 * time.Second
	windowsCommandTimeout                     = 1 * time.Minute
	validateNetworkPolicyTimeout              = 3 * time.Minute
	validateDNSTimeout                        = 2 * time.Minute
	firstMasterRegexStr                       = "^k8s-master-"
	podLookupRetries                          = 5
)

var (
	cfg                             config.Config
	eng                             engine.Engine
	masterSSHPort                   string
	masterSSHPrivateKeyFilepath     string
	longRunningApacheDeploymentName string
	sshConn                         *remote.Connection
	kubeConfig                      *Config
	firstMasterRegexp               *regexp.Regexp
)

var _ = BeforeSuite(func() {
	cwd, _ := os.Getwd()
	rootPath := filepath.Join(cwd, "../../..") // The current working dir of these tests is down a few levels from the root of the project. We should traverse up that path so we can find the _output dir
	c, err := config.ParseConfig()
	c.CurrentWorkingDir = rootPath
	Expect(err).NotTo(HaveOccurred())
	cfg = *c // We have to do this because golang anon functions and scoping and stuff

	engCfg, err := engine.ParseConfig(c.CurrentWorkingDir, c.ClusterDefinition, c.Name)
	Expect(err).NotTo(HaveOccurred())
	csInput, err := engine.ParseInput(engCfg.ClusterDefinitionTemplate)
	Expect(err).NotTo(HaveOccurred())
	isUpdate := cfg.Name != ""
	validate := false
	csGenerated, err := engine.ParseOutput(engCfg.GeneratedDefinitionPath+"/apimodel.json", validate, isUpdate)
	Expect(err).NotTo(HaveOccurred())
	eng = engine.Engine{
		Config:             engCfg,
		ClusterDefinition:  csInput,
		ExpandedDefinition: csGenerated,
	}
	masterNodes, err := node.GetByRegexWithRetry("^k8s-master-", 3*time.Minute, cfg.Timeout)
	Expect(err).NotTo(HaveOccurred())
	masterName := masterNodes[0].Metadata.Name
	if strings.Contains(masterName, "vmss") {
		masterSSHPort = "50001"
	} else {
		masterSSHPort = "22"
	}
	masterSSHPrivateKeyFilepath = cfg.GetSSHKeyPath()
	longRunningApacheDeploymentName = "php-apache-long-running"
	kubeConfig, err = GetConfigWithRetry(3*time.Second, cfg.Timeout)
	Expect(err).NotTo(HaveOccurred())
	sshConn, err = remote.NewConnectionWithRetry(kubeConfig.GetServerName(), masterSSHPort, eng.ExpandedDefinition.Properties.LinuxProfile.AdminUsername, masterSSHPrivateKeyFilepath, 3*time.Second, cfg.Timeout)
	Expect(err).NotTo(HaveOccurred())
	success := false
	for i := 0; i < 3; i++ {
		sshAddErr := util.AddToSSHKeyChain(masterSSHPrivateKeyFilepath)
		if sshAddErr == nil {
			success = true
			break
		}
		if i > 1 {
			log.Printf("Error while setting up ssh key forwarding:%s\n", sshAddErr)
		}
		time.Sleep(10 * time.Second)
	}
	Expect(success).To(BeTrue())
	firstMasterRegexp, err = regexp.Compile(firstMasterRegexStr)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	if cfg.DebugAfterSuite {
		cmd := exec.Command("k", "get", "deployments,pods,svc,daemonsets,configmaps,endpoints,jobs,clusterroles,clusterrolebindings,roles,rolebindings,storageclasses", "--all-namespaces", "-o", "wide")
		out, err := cmd.CombinedOutput()
		log.Printf("%s\n", out)
		if err != nil {
			log.Printf("Error: Unable to print all cluster resources\n")
		}
		pod.PrintPodsLogs("kube-addon-manager", "kube-system")
		pod.PrintPodsLogs("kube-proxy", "kube-system")
		pod.PrintPodsLogs("kube-scheduler", "kube-system")
		pod.PrintPodsLogs("kube-apiserver", "kube-system")
		pod.PrintPodsLogs("kube-controller-manager", "kube-system")
	}
})

var _ = Describe("Azure Container Cluster using the Kubernetes Orchestrator", func() {
	Describe("regardless of agent pool type", func() {
		It("should validate host OS DNS", func() {
			var nodeList *node.List
			var err error
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				nodeList, err = node.GetReady()
			} else {
				var nodes []node.Node
				nodes, err = node.GetByRegexWithRetry(firstMasterRegexStr, 3*time.Minute, cfg.Timeout)
				nodeList = &node.List{
					Nodes: nodes,
				}
			}
			Expect(err).NotTo(HaveOccurred())
			hostOSDNSValidateScript := "host-os-dns-validate.sh"
			err = sshConn.CopyTo(hostOSDNSValidateScript)
			Expect(err).NotTo(HaveOccurred())
			envString := "NODE_HOSTNAMES='"
			for _, node := range nodeList.Nodes {
				envString += fmt.Sprintf("%s ", node.Metadata.Name)
			}
			lookupRetries := 3
			envString += fmt.Sprintf("' LOOKUP_RETRIES=%d", lookupRetries)
			hostOSDNSValidationCommand := fmt.Sprintf("%s /tmp/%s", envString, hostOSDNSValidateScript)
			var success bool
			// Retry for up to 5 minutes host vm DNS validation
			for i := 0; i < 30; i++ {
				err := sshConn.Execute(hostOSDNSValidationCommand, true)
				if err == nil {
					success = true
					break
				} else {
					time.Sleep(10 * time.Second)
				}
			}
			Expect(success).To(BeTrue())
			hostOSDNSValidationCommand = fmt.Sprintf("\"%s /tmp/%s\"", envString, hostOSDNSValidateScript)
			for _, node := range nodeList.Nodes {
				if node.IsLinux() && !firstMasterRegexp.MatchString(node.Metadata.Name) {
					err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+hostOSDNSValidateScript)
					Expect(err).NotTo(HaveOccurred())
					err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, hostOSDNSValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
				}
			}
		})

		It("should have the expected k8s version", func() {
			nodeList, err := node.GetReady()
			Expect(err).NotTo(HaveOccurred())
			for _, node := range nodeList.Nodes {
				Expect("v" + eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(node.Version()))
			}
		})

		It("should display the installed Ubuntu version on the master node", func() {
			if eng.ExpandedDefinition.Properties.MasterProfile.IsUbuntu() {
				lsbReleaseCmd := fmt.Sprintf("lsb_release -a && uname -r")
				err := sshConn.Execute(lsbReleaseCmd, true)
				Expect(err).NotTo(HaveOccurred())
				kernelVerCmd := fmt.Sprintf("cat /proc/version")
				err = sshConn.Execute(kernelVerCmd, true)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("This is not an ubuntu master")
			}
		})

		It("should display the installed docker runtime on all nodes", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.RequiresDocker() {
					nodeList, err := node.GetReady()
					Expect(err).NotTo(HaveOccurred())
					dockerVersionCmd := fmt.Sprintf("\"docker version\"")
					for _, node := range nodeList.Nodes {
						if node.IsWindows() {
							if eng.ExpandedDefinition.Properties.WindowsProfile != nil && !eng.ExpandedDefinition.Properties.WindowsProfile.SSHEnabled {
								log.Printf("Can't ssh into Windows node %s because there is no SSH listener", node.Metadata.Name)
								continue
							}
						}
						err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, dockerVersionCmd, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				} else {
					Skip("Skip docker validations on non-docker-backed clusters")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate that every linux node has a root password", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodeList, err := node.GetReady()
					Expect(err).NotTo(HaveOccurred())
					rootPasswdCmd := fmt.Sprintf("\"sudo grep '^root:[!*]:' /etc/shadow\" && exit 1 || exit 0")
					for _, node := range nodeList.Nodes {
						if node.IsUbuntu() {
							err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, rootPasswdCmd, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
						}
					}
				} else {
					Skip("This config is only available on VHD")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate Ubuntu host OS network configuration on all nodes", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodeList, err := node.GetReady()
					Expect(err).NotTo(HaveOccurred())
					netConfigValidateScript := "net-config-validate.sh"
					err = sshConn.CopyTo(netConfigValidateScript)
					Expect(err).NotTo(HaveOccurred())
					netConfigValidationCommand := fmt.Sprintf("\"/tmp/%s\"", netConfigValidateScript)
					err = sshConn.Execute(netConfigValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, node := range nodeList.Nodes {
						if node.IsUbuntu() && !firstMasterRegexp.MatchString(node.Metadata.Name) {
							err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+netConfigValidateScript)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, netConfigValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
						}
					}
				} else {
					Skip("This config is only available on VHD")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate all CIS VHD-paved files", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodeList, err := node.GetReady()
					Expect(err).NotTo(HaveOccurred())
					CISFilesValidateScript := "CIS-files-validate.sh"
					err = sshConn.CopyTo(CISFilesValidateScript)
					Expect(err).NotTo(HaveOccurred())
					CISValidationCommand := fmt.Sprintf("\"/tmp/%s\"", CISFilesValidateScript)
					err = sshConn.Execute(CISValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, node := range nodeList.Nodes {
						if !firstMasterRegexp.MatchString(node.Metadata.Name) {
							err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+CISFilesValidateScript)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, CISValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
							fmt.Println(err)
						}
					}
				} else {
					Skip("This config is only available on VHD")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate kernel module configuration", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodeList, err := node.GetReady()
					Expect(err).NotTo(HaveOccurred())
					modprobeConfigValidateScript := "modprobe-config-validate.sh"
					err = sshConn.CopyTo(modprobeConfigValidateScript)
					Expect(err).NotTo(HaveOccurred())
					netConfigValidationCommand := fmt.Sprintf("\"/tmp/%s\"", modprobeConfigValidateScript)
					err = sshConn.Execute(netConfigValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, node := range nodeList.Nodes {
						if node.IsUbuntu() && !firstMasterRegexp.MatchString(node.Metadata.Name) {
							err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+modprobeConfigValidateScript)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, netConfigValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
						}
					}
				} else {
					Skip("This config is only available on VHD")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate installed software packages", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				nodeList, err := node.GetReady()
				Expect(err).NotTo(HaveOccurred())
				installedPackagesValidateScript := "ubuntu-installed-packages-validate.sh"
				err = sshConn.CopyTo(installedPackagesValidateScript)
				Expect(err).NotTo(HaveOccurred())
				installedPackagesValidationCommand := fmt.Sprintf("\"/tmp/%s\"", installedPackagesValidateScript)
				err = sshConn.Execute(installedPackagesValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
				for _, node := range nodeList.Nodes {
					if node.IsUbuntu() && !firstMasterRegexp.MatchString(node.Metadata.Name) {
						err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+installedPackagesValidateScript)
						Expect(err).NotTo(HaveOccurred())
						err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, installedPackagesValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate that every linux node has the right sshd config", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodeList, err := node.GetReady()
					Expect(err).NotTo(HaveOccurred())
					sshdConfigValidateScript := "sshd-config-validate.sh"
					err = sshConn.CopyTo(sshdConfigValidateScript)
					Expect(err).NotTo(HaveOccurred())
					sshdConfigValidationCommand := fmt.Sprintf("\"/tmp/%s\"", sshdConfigValidateScript)
					err = sshConn.Execute(sshdConfigValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, node := range nodeList.Nodes {
						if node.IsUbuntu() && !firstMasterRegexp.MatchString(node.Metadata.Name) {
							err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+sshdConfigValidateScript)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, sshdConfigValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
						}
					}
				} else {
					Skip("This config is only available on VHD")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate password enforcement configuration", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodeList, err := node.GetReady()
					Expect(err).NotTo(HaveOccurred())
					pwQualityValidateScript := "pwquality-validate.sh"
					err = sshConn.CopyTo(pwQualityValidateScript)
					Expect(err).NotTo(HaveOccurred())
					pwQualityValidationCommand := fmt.Sprintf("\"/tmp/%s\"", pwQualityValidateScript)
					err = sshConn.Execute(pwQualityValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, node := range nodeList.Nodes {
						if node.IsUbuntu() && !firstMasterRegexp.MatchString(node.Metadata.Name) {
							err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+pwQualityValidateScript)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, pwQualityValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
						}
					}
				} else {
					Skip("This config is only available on VHD")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate auditd configuration", func() {
			var auditDNodePrefixes []string
			var lowPriVMSSPrefixes []string
			if eng.ExpandedDefinition.Properties.MasterProfile != nil {
				if to.Bool(eng.ExpandedDefinition.Properties.MasterProfile.AuditDEnabled) {
					auditDNodePrefixes = append(auditDNodePrefixes, "k8s-master-")
				}
			}
			for _, profile := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
				if profile.IsLowPriorityScaleSet() {
					lowPriVMSSPrefixes = append(lowPriVMSSPrefixes, "k8s-"+profile.Name)
				} else if to.Bool(profile.AuditDEnabled) {
					auditDNodePrefixes = append(auditDNodePrefixes, profile.Name)
				}
			}
			nodeList, err := node.GetReady()
			Expect(err).NotTo(HaveOccurred())
			auditdValidateScript := "auditd-validate.sh"
			err = sshConn.CopyTo(auditdValidateScript)
			Expect(err).NotTo(HaveOccurred())
			for _, node := range nodeList.Nodes {
				if !node.HasSubstring(lowPriVMSSPrefixes) && node.IsUbuntu() {
					var enabled bool
					if node.HasSubstring(auditDNodePrefixes) {
						enabled = true
					}
					err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+auditdValidateScript)
					Expect(err).NotTo(HaveOccurred())
					auditdValidationCommand := fmt.Sprintf("\"ENABLED=%t /tmp/%s\"", enabled, auditdValidateScript)
					err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, auditdValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
				}
			}
		})

		It("should be able to kubectl port-forward to a running pod", func() {
			deploymentNamespace := "default"

			var deploy *deployment.Deployment
			var err error
			var pods []pod.Pod

			testPortForward := func(deploymentName string) {
				running, podWaitErr := pod.WaitOnSuccesses(deploymentName, deploymentNamespace, 3, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(podWaitErr).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				pods, err = deploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(pods)).To(Equal(1))
				for _, p := range pods {
					func() {
						By("Ensuring that the pod is running")
						var running bool
						running, err = p.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						Expect(running).To(Equal(true))
						By("Running kubectl port-forward")
						var proxyCmd *exec.Cmd
						var proxyStdout, proxyStderr io.ReadCloser
						var proxyStdoutReader, proxyStderrReader *bufio.Reader
						success := false
						for i := 0; i < 5; i++ {
							if i > 1 {
								log.Printf("Waiting for retry...\n")
								time.Sleep(10 * time.Second)
							}
							proxyCmd = exec.Command("k", "port-forward", p.Metadata.Name, "8123:80")
							proxyCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
							proxyStdout, err = proxyCmd.StdoutPipe()
							Expect(err).NotTo(HaveOccurred())
							proxyStderr, err = proxyCmd.StderrPipe()
							Expect(err).NotTo(HaveOccurred())
							util.PrintCommand(proxyCmd)
							err = proxyCmd.Start()
							if err != nil {
								log.Printf("kubectl port-forward start error: %v\n", err)
								continue
							}
							proxyStdoutReader = bufio.NewReader(proxyStdout)
							proxyStderrReader = bufio.NewReader(proxyStderr)
							proxyOutStr, outErr := proxyStdoutReader.ReadString('\n')
							log.Printf("kubectl port-forward stdout: %s\n", proxyOutStr)
							if outErr != nil {
								proxyErrStr, _ := proxyStderrReader.ReadString('\n') // returns EOF error, ignore it
								log.Printf("kubectl port-forward stderr: %s\n", proxyErrStr)
								continue
							}
							defer func() {
								syscall.Kill(-proxyCmd.Process.Pid, syscall.SIGKILL)
								_, waiterr := proxyCmd.Process.Wait()
								if waiterr != nil {
									log.Printf("kubectl port-forward - no wait error\n")
								} else {
									log.Printf("kubectl port-forward - wait returned err: %v\n", waiterr)
								}
							}()
							log.Printf("kubectl port-forward running as pid: %d\n", proxyCmd.Process.Pid)
							success = true
							break
						}
						Expect(success).To(Equal(true))
						By("Running curl to access the forwarded port")
						url := fmt.Sprintf("http://%s:%v", "localhost", 8123)
						cmd := exec.Command("curl", "--max-time", "60", "--retry", "10", "--retry-delay", "10", "--retry-max-time", "120", url)
						util.PrintCommand(cmd)
						var out []byte
						out, err = cmd.CombinedOutput()
						log.Printf("%s\n", out)
						Expect(err).NotTo(HaveOccurred())
					}()
				}
			}
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			if eng.AnyAgentIsLinux() {
				By("Creating a Linux nginx deployment")
				deploymentPrefix := "portforwardlinux"
				deploymentName := fmt.Sprintf("%s-%v", deploymentPrefix, r.Intn(9999))
				deploy, err = deployment.CreateLinuxDeployDeleteIfExists(deploymentPrefix, "library/nginx:latest", deploymentName, deploymentNamespace, "")
				Expect(err).NotTo(HaveOccurred())
				testPortForward(deploymentName)
				err = deploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			}
			if eng.HasWindowsAgents() {
				By("Creating a Windows IIS deployment")
				if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.15.0") {
					windowsImages, err := eng.GetWindowsTestImages()
					Expect(err).NotTo(HaveOccurred())
					deploymentPrefix := "portforwardwindows"
					deploymentName := fmt.Sprintf("%s-%v", deploymentPrefix, r.Intn(9999))
					deploy, err = deployment.CreateWindowsDeployDeleteIfExist(deploymentPrefix, windowsImages.IIS, deploymentName, deploymentNamespace, "")
					Expect(err).NotTo(HaveOccurred())
					testPortForward(deploymentName)
					err = deploy.Delete(util.DefaultDeleteRetries)
					Expect(err).NotTo(HaveOccurred())
				} else {
					Skip("kubectl port-forward only works on Windows nodes with Kubernetes 1.15+")
					// Reference: https://github.com/kubernetes/kubernetes/pull/75479
				}
			}
		})

		It("should report all nodes in a Ready state", func() {
			var expectedReadyNodes int
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				expectedReadyNodes = eng.NodeCount()
				log.Printf("Checking for %d Ready nodes\n", expectedReadyNodes)
			} else {
				expectedReadyNodes = -1
			}
			ready := node.WaitOnReady(expectedReadyNodes, 10*time.Second, cfg.Timeout)
			cmd := exec.Command("k", "get", "nodes", "-o", "wide")
			out, _ := cmd.CombinedOutput()
			log.Printf("%s\n", out)
			Expect(ready).To(Equal(true))
		})

		It("should have node labels and annotations", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				totalNodeCount := eng.NodeCount()
				masterNodes, err := node.GetByRegexWithRetry(firstMasterRegexStr, 3*time.Minute, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				nodes := totalNodeCount - len(masterNodes)
				nodeList, err := node.GetByLabel("foo")
				Expect(err).NotTo(HaveOccurred())
				Expect(len(nodeList)).To(Equal(nodes))
				nodeList, err = node.GetByAnnotations("foo", "bar")
				Expect(err).NotTo(HaveOccurred())
				Expect(len(nodeList)).To(Equal(nodes))
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should have node labels specific to masters or agents", func() {
			nodes, err := node.GetWithRetry(1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				Expect(len(nodes)).To(Equal(eng.NodeCount()))
			}
			for _, node := range nodes {
				role := "master"
				if !strings.HasPrefix(node.Metadata.Name, "k8s-master-") {
					if eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
						continue
					} else {
						role = "agent"
					}
				}
				labels := node.Metadata.Labels
				// See https://github.com/Azure/aks-engine/issues/1660
				if node.IsWindows() && common.IsKubernetesVersionGe(
					eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.16.0-alpha.1") {
					Skip("Kubernetes 1.16 on Windows needs node labels applied")
				}
				Expect(labels).To(HaveKeyWithValue("kubernetes.io/role", role))
				Expect(labels).To(HaveKey(fmt.Sprintf("node-role.kubernetes.io/%s", role)))
			}
		})

		It("should print cluster resources", func() {
			cmd := exec.Command("k", "get", "deployments,pods,svc,daemonsets,configmaps,endpoints,jobs,clusterroles,clusterrolebindings,roles,rolebindings,storageclasses", "--all-namespaces", "-o", "wide")
			out, err := cmd.CombinedOutput()
			log.Printf("%s\n", out)
			if err != nil {
				log.Printf("Error: Unable to print all cluster resources\n")
			}
		})

		It("should have DNS pod running", func() {
			var err error
			var running bool
			if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.12.0") {
				By("Ensuring that coredns is running")
				running, err = pod.WaitOnSuccesses("coredns", "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)

			} else {
				By("Ensuring that kube-dns is running")
				running, err = pod.WaitOnSuccesses("kube-dns", "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))
		})

		It("should have core kube-system componentry running", func() {
			coreComponents := []string{"kube-proxy", "kube-addon-manager", "kube-apiserver", "kube-controller-manager", "kube-scheduler"}
			if !common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.13.0") {
				coreComponents = append(coreComponents, "heapster")
			}
			for _, componentName := range coreComponents {
				By(fmt.Sprintf("Ensuring that %s is Running", componentName))
				running, err := pod.WaitOnSuccesses(componentName, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
			}
		})

		It("Should not have any unready or crashing pods right after deployment", func() {
			if eng.Config.DebugCrashingPods {
				By("Checking ready status of each pod in kube-system")
				pods, err := pod.GetAll("kube-system")
				Expect(err).NotTo(HaveOccurred())
				Expect(len(pods.Pods)).ToNot(BeZero())
				for _, currentPod := range pods.Pods {
					log.Printf("Checking %s - ready: %t, restarts: %d", currentPod.Metadata.Name, currentPod.Status.ContainerStatuses[0].Ready, currentPod.Status.ContainerStatuses[0].RestartCount)
					Expect(currentPod.Status.ContainerStatuses[0].Ready).To(BeTrue())
					tooManyRestarts := 5
					if strings.Contains(currentPod.Metadata.Name, "cluster-autoscaler") {
						log.Print("need to investigate cluster-autoscaler restarts!")
						tooManyRestarts = 10
					}
					Expect(currentPod.Status.ContainerStatuses[0].RestartCount).To(BeNumerically("<", tooManyRestarts))
				}
			} else {
				Skip("Skipping this DEBUG test")
			}
		})

		It("should have the correct IP address for the apiserver", func() {
			pods, err := pod.GetAllRunningByPrefixWithRetry("kube-apiserver", "kube-system", 3*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			By("Ensuring that the correct IP address has been applied to the apiserver")
			expectedIPAddress := eng.ExpandedDefinition.Properties.MasterProfile.FirstConsecutiveStaticIP
			if eng.ExpandedDefinition.Properties.MasterProfile.HasMultipleNodes() {
				firstMasterIP := net.ParseIP(eng.ExpandedDefinition.Properties.MasterProfile.FirstConsecutiveStaticIP).To4()
				expectedIP := net.IP{firstMasterIP[0], firstMasterIP[1], firstMasterIP[2], firstMasterIP[3] + byte(common.DefaultInternalLbStaticIPOffset)}
				if eng.ExpandedDefinition.Properties.MasterProfile.IsVirtualMachineScaleSets() {
					expectedIP = net.IP{firstMasterIP[0], firstMasterIP[1], byte(255), byte(common.DefaultInternalLbStaticIPOffset)}
				}
				expectedIPAddress = expectedIP.String()
			}

			actualIPAddress, err := pods[0].Spec.Containers[0].GetArg("--advertise-address")
			Expect(err).NotTo(HaveOccurred())
			Expect(actualIPAddress).To(Equal(expectedIPAddress))
		})

		It("should have addons running", func() {
			for _, addonName := range []string{"tiller", "aci-connector", "cluster-autoscaler", "blobfuse-flexvolume", "smb-flexvolume", "keyvault-flexvolume", "kubernetes-dashboard", "rescheduler", "metrics-server", "nvidia-device-plugin", "container-monitoring", "azure-cni-networkmonitor", "azure-npm-daemonset", "ip-masq-agent"} {
				var addonPods = []string{addonName}
				var addonNamespace = "kube-system"
				switch addonName {
				case "blobfuse-flexvolume":
					addonPods = []string{"blobfuse-flexvol-installer"}
				case "smb-flexvolume":
					addonPods = []string{"smb-flexvol-installer"}
				case "container-monitoring":
					addonPods = []string{"omsagent"}
				case "azure-npm-daemonset":
					addonPods = []string{"azure-npm"}
				}
				if hasAddon, addon := eng.HasAddon(addonName); hasAddon {
					for _, addonPod := range addonPods {
						By(fmt.Sprintf("Ensuring that the %s addon is Running", addonName))
						running, err := pod.WaitOnSuccesses(addonPod, addonNamespace, kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						Expect(running).To(Equal(true))
						By(fmt.Sprintf("Ensuring that the correct resources have been applied for %s", addonPod))
						pods, err := pod.GetAllRunningByPrefixWithRetry(addonPod, addonNamespace, 3*time.Second, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						for i, c := range addon.Containers {
							pod := pods[0]
							container := pod.Spec.Containers[i]
							err := container.ValidateResources(c)
							Expect(err).NotTo(HaveOccurred())
						}
					}
				} else {
					fmt.Printf("%s disabled for this cluster, will not test\n", addonName)
				}
			}
		})

		It("should have the correct tiller configuration", func() {
			if hasTiller, tillerAddon := eng.HasAddon("tiller"); hasTiller {
				running, err := pod.WaitOnSuccesses("tiller", "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				pods, err := pod.GetAllRunningByPrefixWithRetry("tiller-deploy", "kube-system", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				By("Ensuring that the correct max-history has been applied")
				maxHistory := tillerAddon.Config["max-history"]
				// There is only one tiller pod and one container in that pod
				actualTillerMaxHistory, err := pods[0].Spec.Containers[0].GetEnvironmentVariable("TILLER_HISTORY_MAX")
				Expect(err).NotTo(HaveOccurred())
				Expect(actualTillerMaxHistory).To(Equal(maxHistory))
			} else {
				Skip("tiller disabled for this cluster, will not test")
			}
		})

		It("should have the expected omsagent cluster footprint", func() {
			if hasContainerMonitoring, _ := eng.HasAddon("container-monitoring"); hasContainerMonitoring {
				By("Validating the omsagent replicaset")
				running, err := pod.WaitOnSuccesses("omsagent-rs", "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				pods, err := pod.GetAllRunningByPrefixWithRetry("omsagent-rs", "kube-system", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				By("Ensuring that the kubepodinventory plugin is writing data successfully")
				pass, err := pods[0].ValidateOmsAgentLogs("kubePodInventoryEmitStreamSuccess", 1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())
				By("Ensuring that the kubenodeinventory plugin is writing data successfully")
				pass, err = pods[0].ValidateOmsAgentLogs("kubeNodeInventoryEmitStreamSuccess", 1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())
				By("Validating the omsagent daemonset")
				running, err = pod.WaitOnSuccesses("omsagent", "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				pods, err = pod.GetAllRunningByPrefixWithRetry("omsagent", "kube-system", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				By("Ensuring that the cadvisor_perf plugin is writing data successfully")
				pass, err = pods[0].ValidateOmsAgentLogs("cAdvisorPerfEmitStreamSuccess", 1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())
				By("Ensuring that the containerinventory plugin is writing data successfully")
				pass, err = pods[0].ValidateOmsAgentLogs("containerInventoryEmitStreamSuccess", 1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())
			} else {
				Skip("container monitoring disabled for this cluster, will not test")
			}
		})

		It("should be able to launch a long-running container networking DNS liveness pod", func() {
			p, err := pod.CreatePodFromFileIfNotExist(filepath.Join(WorkloadDir, "dns-liveness.yaml"), "dns-liveness", "default", 1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			running, err := p.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))
		})

		It("should be able to launch a long running HTTP listener and svc endpoint", func() {
			By("Creating a php-apache deployment")
			phpApacheDeploy, err := deployment.CreateLinuxDeployIfNotExist("deis/hpa-example", longRunningApacheDeploymentName, "default", "--requests=cpu=10m,memory=10M")
			Expect(err).NotTo(HaveOccurred())

			By("Ensuring that php-apache pod is running")
			running, err := pod.WaitOnSuccesses(longRunningApacheDeploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))

			By("Ensuring that the php-apache pod has outbound internet access")
			pods, err := phpApacheDeploy.PodsRunning()
			Expect(err).NotTo(HaveOccurred())
			for _, p := range pods {
				pass, outboundErr := p.CheckLinuxOutboundConnection(5*time.Second, cfg.Timeout)
				Expect(outboundErr).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())
			}

			By("Exposing TCP 80 internally on the php-apache deployment")
			err = phpApacheDeploy.ExposeIfNotExist("ClusterIP", 80, 80)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have stable external container networking as we recycle a bunch of pods", func() {
			name := fmt.Sprintf("alpine-%s", cfg.Name)
			command := fmt.Sprintf("nc -vz 8.8.8.8 53 || nc -vz 8.8.4.4 53")
			successes, err := pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, retryCommandsTimeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
		})

		It("should have stable internal container networking as we recycle a bunch of pods", func() {
			name := fmt.Sprintf("alpine-%s", cfg.Name)
			var command string
			if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.12.0") {
				command = fmt.Sprintf("nc -vz kubernetes 443 && nc -vz kubernetes.default.svc 443 && nc -vz kubernetes.default.svc.cluster.local 443")
			} else {
				command = fmt.Sprintf("nc -vz kubernetes 443")
			}
			successes, err := pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, retryCommandsTimeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
		})

		It("should have stable pod-to-pod networking", func() {
			if eng.AnyAgentIsLinux() {
				By("Creating a test php-apache deployment")
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				By("Creating another pod that will connect to the php-apache pod")
				commandString := fmt.Sprintf("nc -vz %s.default.svc.cluster.local 80", longRunningApacheDeploymentName)
				consumerPodName := fmt.Sprintf("consumer-pod-%s-%v", cfg.Name, r.Intn(99999))
				successes, err := pod.RunCommandMultipleTimes(pod.RunLinuxPod, "busybox", consumerPodName, commandString, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, retryCommandsTimeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))
			} else {
				Skip("Pod-to-pod network tests only valid on Linux clusters")
			}
		})

		It("should have functional container networking DNS", func() {
			By("Ensuring that we have functional DNS resolution from a linux container")
			j, err := job.CreateJobFromFileDeleteIfExists(filepath.Join(WorkloadDir, "validate-dns-linux.yaml"), "validate-dns-linux", "default", 3*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			ready, err := j.WaitOnSucceeded(sleepBetweenRetriesWhenWaitingForPodReady, validateDNSTimeout)
			delErr := j.Delete(util.DefaultDeleteRetries)
			if delErr != nil {
				fmt.Printf("could not delete job %s\n", j.Metadata.Name)
				fmt.Println(delErr)
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(ready).To(Equal(true))

			if eng.HasWindowsAgents() {
				By("Ensuring that we have functional DNS resolution from a windows container")
				windowsImages, imgErr := eng.GetWindowsTestImages()
				Expect(imgErr).NotTo(HaveOccurred())
				j, err = job.CreateWindowsJobFromTemplateDeleteIfExists(filepath.Join(WorkloadDir, "validate-dns-windows.yaml"), "validate-dns-windows", "default", windowsImages, 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				ready, err = j.WaitOnSucceeded(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				delErr = j.Delete(util.DefaultDeleteRetries)
				if delErr != nil {
					fmt.Printf("could not delete job %s\n", j.Metadata.Name)
					fmt.Println(delErr)
				}
				Expect(err).NotTo(HaveOccurred())
				Expect(ready).To(Equal(true))
			}

			By("Ensuring that we have stable external DNS resolution as we recycle a bunch of pods")
			name := fmt.Sprintf("alpine-%s", cfg.Name)
			command := fmt.Sprintf("nc -vz bbc.co.uk 80 || nc -vz google.com 443 || nc -vz microsoft.com 80")
			successes, err := pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, retryCommandsTimeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
		})

		It("should be able to access the dashboard", func() {
			if hasDashboard, _ := eng.HasAddon("kubernetes-dashboard"); hasDashboard {
				By("Ensuring that the kubernetes-dashboard service is Running")
				s, err := service.Get("kubernetes-dashboard", "kube-system")
				Expect(err).NotTo(HaveOccurred())
				By("Ensuring that we can connect via HTTPS to the dashboard on any one node")
				dashboardPort := 443
				port := s.GetNodePort(dashboardPort)
				nodeList, err := node.GetReady()
				Expect(err).NotTo(HaveOccurred())
				var success bool
				for _, node := range nodeList.Nodes {
					if success {
						break
					}
					if node.IsLinux() {
						// Allow 3 retries for each node
						for i := 0; i < 3; i++ {
							address := node.Status.GetAddressByType("InternalIP")
							if address == nil {
								log.Printf("One of our nodes does not have an InternalIP value!: %s\n", node.Metadata.Name)
							}
							Expect(address).NotTo(BeNil())
							dashboardURL := fmt.Sprintf("http://%s:%v", address.Address, port)
							curlCMD := fmt.Sprintf("curl --max-time 60 %s", dashboardURL)
							err := sshConn.Execute(curlCMD, false)
							if err == nil {
								success = true
								break
							}
							time.Sleep(1 * time.Second)
						}
					}
				}
				Expect(success).To(BeTrue())
			} else {
				Skip("kubernetes-dashboard disabled for this cluster, will not test")
			}
		})
	})

	Describe("with a windows agent pool", func() {
		It("kubelet service should be able to recover when the docker service is stopped", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.HasWindowsAgents() {
					if eng.ExpandedDefinition.Properties.WindowsProfile != nil && eng.ExpandedDefinition.Properties.WindowsProfile.SSHEnabled {
						nodeList, err := node.GetReady()
						Expect(err).NotTo(HaveOccurred())
						simulateDockerdCrashScript := "simulate-dockerd-crash.cmd"
						err = sshConn.CopyTo(simulateDockerdCrashScript)
						Expect(err).NotTo(HaveOccurred())
						for _, node := range nodeList.Nodes {
							if node.IsWindows() {
								By(fmt.Sprintf("simulating docker and subsequent kubelet service crash on node: %s", node.Metadata.Name))
								err = sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+simulateDockerdCrashScript)
								Expect(err).NotTo(HaveOccurred())
								simulateDockerCrashCommand := fmt.Sprintf("\"/tmp/%s\"", simulateDockerdCrashScript)
								err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, simulateDockerCrashCommand, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
								Expect(err).NotTo(HaveOccurred())
							}
						}

						log.Print("Waiting 1 minute to allow nodes to report not ready state after the crash occurred\n")
						time.Sleep(1 * time.Minute)

						for _, node := range nodeList.Nodes {
							if node.IsWindows() {
								By(fmt.Sprintf("restarting kubelet service on node: %s", node.Metadata.Name))
								restartKubeletCommand := fmt.Sprintf("\"Powershell Start-Service kubelet\"")
								err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, restartKubeletCommand, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
								Expect(err).NotTo(HaveOccurred())
							}
						}

						var expectedReadyNodes int
						if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
							expectedReadyNodes = len(nodeList.Nodes)
							log.Printf("Checking for %d Ready nodes\n", expectedReadyNodes)
						} else {
							expectedReadyNodes = -1
						}
						ready := node.WaitOnReady(expectedReadyNodes, 1*time.Minute, cfg.Timeout)
						cmd2 := exec.Command("k", "get", "nodes", "-o", "wide")
						out2, _ := cmd2.CombinedOutput()
						log.Printf("%s\n", out2)
						if !ready {
							log.Printf("Error: Not all nodes in a healthy state\n")
						}
						Expect(ready).To(Equal(true))
					} else {
						Skip("Windows SSH tests only work if WindowsProfile.SSHEnabled is true")
					}
				} else {
					Skip("Docker service recovery test is Windows only")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})
	})

	Describe("with a linux agent pool", func() {
		It("should be able to produce working LoadBalancers", func() {
			if eng.AnyAgentIsLinux() {
				By("Creating a nginx deployment")
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				serviceName := "ingress-nginx"
				deploymentPrefix := fmt.Sprintf("%s-%s", serviceName, cfg.Name)
				deploymentName := fmt.Sprintf("%s-%v", deploymentPrefix, r.Intn(99999))
				deploy, err := deployment.CreateLinuxDeployDeleteIfExists(deploymentPrefix, "library/nginx:latest", deploymentName, "default", "--labels=app="+serviceName)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can create an ILB service attachment")
				sILB, err := service.CreateServiceFromFileDeleteIfExist(filepath.Join(WorkloadDir, "ingress-nginx-ilb.yaml"), serviceName+"-ilb", "default")
				Expect(err).NotTo(HaveOccurred())
				err = sILB.WaitForIngress(cfg.Timeout, 5*time.Second)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can create a curl pod to connect to the service")
				deploymentPrefix = fmt.Sprintf("ilb-test-curl-deployment")
				curlDeploymentName := fmt.Sprintf("%s-%v", deploymentPrefix, r.Intn(99999))
				curlDeploy, err := deployment.CreateLinuxDeployDeleteIfExists(deploymentPrefix, "library/nginx:latest", curlDeploymentName, "default", "--replicas=2")
				Expect(err).NotTo(HaveOccurred())
				running, err := pod.WaitOnSuccesses(curlDeploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				curlPods, err := curlDeploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				By("Ensuring we can connect to the ILB service from another pod")
				var success bool
				for _, curlPod := range curlPods {
					pass, curlErr := curlPod.ValidateCurlConnection(sILB.Status.LoadBalancer.Ingress[0]["ip"], 30*time.Second, 3*time.Minute)
					if curlErr == nil && pass {
						success = true
						break
					} else {
						e := sILB.Describe()
						if e != nil {
							log.Printf("Unable to describe service\n: %s", e)
						}
					}

				}
				Expect(success).To(BeTrue())
				By("Ensuring we can create an ELB service attachment")
				sELB, err := service.CreateServiceFromFileDeleteIfExist(filepath.Join(WorkloadDir, "ingress-nginx-elb.yaml"), serviceName+"-elb", "default")
				Expect(err).NotTo(HaveOccurred())
				err = sELB.WaitForIngress(cfg.Timeout, 5*time.Second)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can connect to the ELB service on the service IP")
				err = sELB.ValidateWithRetry("(Welcome to nginx)", 30*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				By("Ensuring we can connect to the ELB service from another pod")
				success = false
				for _, curlPod := range curlPods {
					pass, curlErr := curlPod.ValidateCurlConnection(sELB.Status.LoadBalancer.Ingress[0]["ip"], 30*time.Second, 3*time.Minute)
					if curlErr == nil && pass {
						success = true
						break
					} else {
						e := sELB.Describe()
						if e != nil {
							log.Printf("Unable to describe service\n: %s", e)
						}
					}

				}
				Expect(success).To(BeTrue())
				err = sILB.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = sELB.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = curlDeploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = deploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("No linux agent was provisioned for this Cluster Definition")
			}
		})

		It("should be able to get nodes metrics", func() {
			if eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.IsRBACEnabled() {
				success := false
				var err error
				var out []byte
				// TODO make a 1st class go func retry implementation of this
				for i := 0; i < 10; i++ {
					cmd := exec.Command("k", "top", "nodes")
					util.PrintCommand(cmd)
					out, err = cmd.CombinedOutput()
					if err == nil {
						success = true
						break
					}
					time.Sleep(1 * time.Minute)
				}
				if err != nil {
					pod.PrintPodsLogs("metrics-server", "kube-system")
					log.Println(string(out))
				}
				Expect(success).To(BeTrue())
			}
		})

		It("should be able to autoscale", func() {
			if eng.AnyAgentIsLinux() && eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs {
				// Inspired by http://blog.kubernetes.io/2016/07/autoscaling-in-kubernetes.html
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				By("Creating a php-apache deployment")
				phpApacheDeploy, err := deployment.CreateLinuxDeployIfNotExist("deis/hpa-example", longRunningApacheDeploymentName, "default", "--requests=cpu=10m,memory=10M")
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring that the php-apache pod is running")
				running, err := pod.WaitOnSuccesses(longRunningApacheDeploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Ensuring that the php-apache pod has outbound internet access")
				pods, err := phpApacheDeploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				for _, p := range pods {
					pass, outboundErr := p.CheckLinuxOutboundConnection(5*time.Second, cfg.Timeout)
					Expect(outboundErr).NotTo(HaveOccurred())
					Expect(pass).To(BeTrue())
				}

				By("Exposing TCP 80 internally on the php-apache deployment")
				err = phpApacheDeploy.ExposeIfNotExist("ClusterIP", 80, 80)
				Expect(err).NotTo(HaveOccurred())

				By("Assigning hpa configuration to the php-apache deployment")
				// Apply autoscale characteristics to deployment
				err = phpApacheDeploy.CreateDeploymentHPADeleteIfExist(80, 1, 10)
				Expect(err).NotTo(HaveOccurred())
				h, err := hpa.Get(longRunningApacheDeploymentName, "default", 10)
				Expect(err).NotTo(HaveOccurred())

				By("Sending load to the php-apache service by creating a 3 replica deployment")
				// Launch a simple busybox pod that wget's continuously to the apache serviceto simulate load
				commandString := fmt.Sprintf("while true; do wget -q -O- http://%s.default.svc.cluster.local; done", longRunningApacheDeploymentName)
				loadTestPrefix := fmt.Sprintf("load-test-%s", cfg.Name)
				loadTestName := fmt.Sprintf("%s-%v", loadTestPrefix, r.Intn(99999))
				numLoadTestPods := 3
				loadTestDeploy, err := deployment.RunLinuxDeployDeleteIfExists(loadTestPrefix, "busybox", loadTestName, "default", commandString, numLoadTestPods)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring there are 3 load test pods")
				running, err = pod.WaitOnSuccesses(loadTestName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				// We should have three load tester pods running
				loadTestPods, err := pod.GetAllRunningByPrefixWithRetry(loadTestPrefix, "default", 5*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(loadTestPods)).To(Equal(numLoadTestPods))

				By("Ensuring we have more than 1 apache-php pods due to hpa enforcement")
				_, err = phpApacheDeploy.WaitForReplicas(2, -1, 5*time.Second, cfg.Timeout)
				if err != nil {
					e := h.Describe()
					Expect(e).NotTo(HaveOccurred())
				}
				Expect(err).NotTo(HaveOccurred())

				By("Stopping load")
				err = loadTestDeploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we only have 1 apache-php pod after stopping load")
				_, err = phpApacheDeploy.WaitForReplicas(-1, 1, 5*time.Second, 20*time.Minute)
				if err != nil {
					e := h.Describe()
					Expect(e).NotTo(HaveOccurred())
				}
				Expect(err).NotTo(HaveOccurred())

				By("Deleting HPA configuration")
				err = h.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("This flavor/version of Kubernetes doesn't support hpa autoscale")
			}
		})

		It("should be able to schedule a pod to a master node", func() {
			By("Creating a pod with master nodeSelector")
			p, err := pod.CreatePodFromFile(filepath.Join(WorkloadDir, "nginx-master.yaml"), "nginx-master", "default", 1*time.Second, cfg.Timeout)
			if err != nil {
				p, err = pod.Get("nginx-master", "default", podLookupRetries)
				Expect(err).NotTo(HaveOccurred())
			}
			running, err := p.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))

			By("validating that master-scheduled pod has outbound internet connectivity")
			pass, err := p.CheckLinuxOutboundConnection(5*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(pass).To(BeTrue())

			By("Cleaning up after ourselves")
			err = p.Delete(util.DefaultDeleteRetries)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("with a GPU-enabled agent pool", func() {
		It("should be able to run a nvidia-gpu job", func() {
			if eng.ExpandedDefinition.Properties.HasNSeriesSKU() {
				version := common.RationalizeReleaseAndVersion(
					common.Kubernetes,
					eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorRelease,
					eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorVersion,
					false,
					eng.HasWindowsAgents())
				if common.IsKubernetesVersionGe(version, "1.10.0") {
					j, err := job.CreateJobFromFileDeleteIfExists(filepath.Join(WorkloadDir, "cuda-vector-add.yaml"), "cuda-vector-add", "default", 3*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					ready, err := j.WaitOnSucceeded(30*time.Second, cfg.Timeout)
					delErr := j.Delete(util.DefaultDeleteRetries)
					if delErr != nil {
						fmt.Printf("could not delete job %s\n", j.Metadata.Name)
						fmt.Println(delErr)
					}
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))
				} else {
					j, err := job.CreateJobFromFileDeleteIfExists(filepath.Join(WorkloadDir, "nvidia-smi.yaml"), "nvidia-smi", "default", 3*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					ready, err := j.WaitOnSucceeded(30*time.Second, cfg.Timeout)
					delErr := j.Delete(util.DefaultDeleteRetries)
					if delErr != nil {
						fmt.Printf("could not delete job %s\n", j.Metadata.Name)
						fmt.Println(delErr)
					}
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))
				}
			} else {
				Skip("This is not a GPU-enabled cluster")
			}
		})
	})

	Describe("with a DC-series SKU agent pool", func() {
		It("should be able to run an SGX job", func() {
			if eng.ExpandedDefinition.Properties.HasDCSeriesSKU() {
				j, err := job.CreateJobFromFileDeleteIfExists(filepath.Join(WorkloadDir, "sgx-test.yaml"), "sgx-test", "default", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				ready, err := j.WaitOnSucceeded(30*time.Second, cfg.Timeout)
				delErr := j.Delete(util.DefaultDeleteRetries)
				if delErr != nil {
					fmt.Printf("could not delete job %s\n", j.Metadata.Name)
					fmt.Println(delErr)
				}
				Expect(err).NotTo(HaveOccurred())
				Expect(ready).To(Equal(true))
			} else {
				Skip("This cluster does not have a DC-series SKU agent pool")
			}
		})
	})

	Describe("with zoned master profile", func() {
		It("should be labeled with zones for each masternode", func() {
			if eng.ExpandedDefinition.Properties.MasterProfile.HasAvailabilityZones() {
				nodes, err := node.GetWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				for _, node := range nodes {
					role := node.Metadata.Labels["kubernetes.io/role"]
					if role == "master" {
						By("Ensuring that we get zones for each master node")
						zones := node.Metadata.Labels["failure-domain.beta.kubernetes.io/zone"]
						contains := strings.Contains(zones, "-")
						Expect(contains).To(Equal(true))
					}
				}
			} else {
				Skip("Availability zones was not configured for master profile for this Cluster Definition")
			}
		})
	})

	Describe("with all zoned agent pools", func() {
		It("should be labeled with zones for each node", func() {
			if eng.ExpandedDefinition.Properties.HasZonesForAllAgentPools() {
				nodes, err := node.GetWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				for _, node := range nodes {
					role := node.Metadata.Labels["kubernetes.io/role"]
					if role == "agent" {
						By("Ensuring that we get zones for each agent node")
						zones := node.Metadata.Labels["failure-domain.beta.kubernetes.io/zone"]
						contains := strings.Contains(zones, "-")
						Expect(contains).To(Equal(true))
					}
				}
			} else {
				Skip("Availability zones was not configured for this Cluster Definition")
			}
		})

		It("should create pv with zone labels and node affinity", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.HasZonesForAllAgentPools() {
					By("Creating a persistent volume claim")
					pvcName := "azure-managed-disk" // should be the same as in pvc-standard.yaml
					pvc, err := persistentvolumeclaims.CreatePersistentVolumeClaimsFromFile(filepath.Join(WorkloadDir, "pvc-standard.yaml"), pvcName, "default")
					Expect(err).NotTo(HaveOccurred())
					ready, err := pvc.WaitOnReady("default", 5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))

					pvList, err := persistentvolume.Get()
					Expect(err).NotTo(HaveOccurred())
					pvZone := ""
					for _, pv := range pvList.PersistentVolumes {
						By("Ensuring that we get zones for the pv")
						// zone is chosen by round-robin across all zones
						pvZone = pv.Metadata.Labels["failure-domain.beta.kubernetes.io/zone"]
						fmt.Printf("pvZone: %s\n", pvZone)
						contains := strings.Contains(pvZone, "-")
						Expect(contains).To(Equal(true))
						// VolumeScheduling feature gate is set to true by default starting v1.10+
						for _, expression := range pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions {
							if expression.Key == "failure-domain.beta.kubernetes.io/zone" {
								By("Ensuring that we get nodeAffinity for each pv")
								value := expression.Values[0]
								fmt.Printf("NodeAffinity value: %s\n", value)
								contains := strings.Contains(value, "-")
								Expect(contains).To(Equal(true))
							}
						}
					}

					By("Launching a pod using the volume claim")
					podName := "zone-pv-pod" // should be the same as in pod-pvc.yaml
					testPod, err := pod.CreatePodFromFile(filepath.Join(WorkloadDir, "pod-pvc.yaml"), podName, "default", 1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					ready, err = testPod.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))

					By("Checking that the pod can access volume")
					valid, err := testPod.ValidatePVC("/mnt/azure", 10, 10*time.Second)
					Expect(valid).To(BeTrue())
					Expect(err).NotTo(HaveOccurred())

					By("Ensuring that attached volume pv has the same zone as the zone of the node")
					nodeName := testPod.Spec.NodeName
					nodeList, err := node.GetByRegexWithRetry(nodeName, 3*time.Minute, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					nodeZone := nodeList[0].Metadata.Labels["failure-domain.beta.kubernetes.io/zone"]
					fmt.Printf("pvZone: %s\n", pvZone)
					fmt.Printf("nodeZone: %s\n", nodeZone)
					Expect(nodeZone == pvZone).To(Equal(true))

					By("Cleaning up after ourselves")
					err = testPod.Delete(util.DefaultDeleteRetries)
					Expect(err).NotTo(HaveOccurred())
					err = pvc.Delete(util.DefaultDeleteRetries)
					Expect(err).NotTo(HaveOccurred())
				} else {
					Skip("Availability zones was not configured for this Cluster Definition")
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})
	})

	Describe("with NetworkPolicy enabled", func() {
		It("should apply various network policies and enforce access to nginx pod", func() {
			if eng.HasNetworkPolicy("calico") || eng.HasNetworkPolicy("azure") || eng.HasNetworkPolicy("cilium") {
				nsDev, nsProd := "development", "production"
				By("Creating development namespace")
				namespaceDev, err := namespace.CreateIfNotExist(nsDev)
				Expect(err).NotTo(HaveOccurred())
				By("Creating production namespace")
				namespaceProd, err := namespace.CreateIfNotExist(nsProd)
				Expect(err).NotTo(HaveOccurred())
				By("Labelling development namespace")
				err = namespaceDev.Label("purpose=development")
				Expect(err).NotTo(HaveOccurred())
				By("Labelling production namespace")
				err = namespaceProd.Label("purpose=production")
				Expect(err).NotTo(HaveOccurred())
				By("Creating frontendProd, backend and network-policy pod deployments")
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				randInt := r.Intn(99999)
				frontendProdDeploymentName := fmt.Sprintf("frontend-prod-%s-%v", cfg.Name, randInt)
				frontendProdDeployment, err := deployment.CreateLinuxDeploy("library/nginx:latest", frontendProdDeploymentName, nsProd, "--labels=app=webapp,role=frontend")
				Expect(err).NotTo(HaveOccurred())
				frontendDevDeploymentName := fmt.Sprintf("frontend-dev-%s-%v", cfg.Name, randInt+100000)
				frontendDevDeployment, err := deployment.CreateLinuxDeploy("library/nginx:latest", frontendDevDeploymentName, nsDev, "--labels=app=webapp,role=frontend")
				Expect(err).NotTo(HaveOccurred())
				backendDeploymentName := fmt.Sprintf("backend-%s-%v", cfg.Name, randInt+200000)
				backendDeployment, err := deployment.CreateLinuxDeploy("library/nginx:latest", backendDeploymentName, nsDev, "--labels=app=webapp,role=backend")
				Expect(err).NotTo(HaveOccurred())
				nwpolicyDeploymentName := fmt.Sprintf("network-policy-%s-%v", cfg.Name, randInt+300000)
				nwpolicyDeployment, err := deployment.CreateLinuxDeploy("library/nginx:latest", nwpolicyDeploymentName, nsDev, "")
				Expect(err).NotTo(HaveOccurred())

				By("Ensure there is a running frontend-prod pod")
				running, err := pod.WaitOnSuccesses(frontendProdDeploymentName, nsProd, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Ensure there is a running frontend-dev pod")
				running, err = pod.WaitOnSuccesses(frontendDevDeploymentName, nsDev, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Ensure there is a running backend pod")
				running, err = pod.WaitOnSuccesses(backendDeploymentName, nsDev, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Ensure there is a running network-policy pod")
				running, err = pod.WaitOnSuccesses(nwpolicyDeploymentName, nsDev, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Ensuring we have outbound internet access from the frontend-prod pods")
				frontendProdPods, err := frontendProdDeployment.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(frontendProdPods)).ToNot(BeZero())
				pl := pod.List{Pods: frontendProdPods}
				pass, err := pl.CheckOutboundConnection(5*time.Second, cfg.Timeout, api.Linux)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())

				By("Ensuring we have outbound internet access from the frontend-dev pods")
				frontendDevPods, err := frontendDevDeployment.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(frontendDevPods)).ToNot(BeZero())
				pl = pod.List{Pods: frontendDevPods}
				pass, err = pl.CheckOutboundConnection(5*time.Second, cfg.Timeout, api.Linux)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())

				By("Ensuring we have outbound internet access from the backend pods")
				backendPods, err := backendDeployment.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(backendPods)).ToNot(BeZero())
				pl = pod.List{Pods: backendPods}
				pass, err = pl.CheckOutboundConnection(5*time.Second, cfg.Timeout, api.Linux)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())

				By("Ensuring we have outbound internet access from the network-policy pods")
				nwpolicyPods, err := nwpolicyDeployment.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(nwpolicyPods)).ToNot(BeZero())
				pl = pod.List{Pods: nwpolicyPods}
				pass, err = pl.CheckOutboundConnection(5*time.Second, cfg.Timeout, api.Linux)
				Expect(err).NotTo(HaveOccurred())
				Expect(pass).To(BeTrue())

				By("Ensuring we have connectivity from network-policy pods to frontend-prod pods")
				pl = pod.List{Pods: nwpolicyPods}
				for _, frontendProdPod := range frontendProdPods {
					pass, err = pl.ValidateCurlConnection(frontendProdPod.Status.PodIP, 30*time.Second, cfg.Timeout)
					if err != nil {
						e := frontendProdPod.Describe()
						if e != nil {
							log.Printf("Unable to describe pod %s\n: %s", frontendProdPod.Metadata.Name, e)
						}
					}
					Expect(err).NotTo(HaveOccurred())
					Expect(pass).To(BeTrue())
				}

				By("Ensuring we have connectivity from network-policy pods to backend pods")
				for _, backendPod := range backendPods {
					pass, err = pl.ValidateCurlConnection(backendPod.Status.PodIP, 30*time.Second, cfg.Timeout)
					if err != nil {
						e := backendPod.Describe()
						if e != nil {
							log.Printf("Unable to describe pod %s\n: %s", backendPod.Metadata.Name, e)
						}
					}
					Expect(err).NotTo(HaveOccurred())
					Expect(pass).To(BeTrue())
				}

				By("Applying a network policy to deny ingress access to app: webapp, role: backend pods in development namespace")
				nwpolicyName, namespace := "backend-deny-ingress", nsDev
				err = networkpolicy.CreateNetworkPolicyFromFile(filepath.Join(PolicyDir, "backend-policy-deny-ingress.yaml"), nwpolicyName, namespace)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we no longer have ingress access from the network-policy pods to backend pods")
				for _, backendPod := range backendPods {
					pass, err = pl.ValidateCurlConnection(backendPod.Status.PodIP, 30*time.Second, validateNetworkPolicyTimeout)
					if err != nil {
						e := backendPod.Describe()
						if e != nil {
							log.Printf("Unable to describe pod %s\n: %s", backendPod.Metadata.Name, e)
						}
					}
					Expect(err).Should(HaveOccurred())
					Expect(pass).To(BeFalse())
				}

				By("Cleaning up after ourselves")
				networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)

				if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.11.0") {
					By("Applying a network policy to only allow ingress access to app: webapp, role: backend pods in development namespace from pods in any namespace with the same labels")
					nwpolicyName, namespace = "backend-allow-ingress-pod-label", nsDev
					err = networkpolicy.CreateNetworkPolicyFromFile(filepath.Join(PolicyDir, "backend-policy-allow-ingress-pod-label.yaml"), nwpolicyName, namespace)
					Expect(err).NotTo(HaveOccurred())

					By("Ensuring we have ingress access from pods with matching labels")
					pl = pod.List{Pods: backendPods}
					for _, backendDstPod := range backendPods {
						pass, err = pl.ValidateCurlConnection(backendDstPod.Status.PodIP, 30*time.Second, cfg.Timeout)
						if err != nil {
							e := backendDstPod.Describe()
							if e != nil {
								log.Printf("Unable to describe pod %s\n: %s", backendDstPod.Metadata.Name, e)
							}
						}
						Expect(err).NotTo(HaveOccurred())
						Expect(pass).To(BeTrue())
					}

					By("Ensuring we don't have ingress access from pods without matching labels")
					pl = pod.List{Pods: nwpolicyPods}
					for _, backendPod := range backendPods {
						pass, err = pl.ValidateCurlConnection(backendPod.Status.PodIP, 30*time.Second, validateNetworkPolicyTimeout)
						if err != nil {
							e := backendPod.Describe()
							if e != nil {
								log.Printf("Unable to describe pod %s\n: %s", backendPod.Metadata.Name, e)
							}
						}
						Expect(err).Should(HaveOccurred())
						Expect(pass).To(BeFalse())
					}

					By("Cleaning up after ourselves")
					networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)

					By("Applying a network policy to only allow ingress access to app: webapp role:backends in development namespace from pods with label app:webapp, role: frontendProd within namespace with label purpose: development")
					nwpolicyName, namespace = "backend-policy-allow-ingress-pod-namespace-label", nsDev
					err = networkpolicy.CreateNetworkPolicyFromFile(filepath.Join(PolicyDir, "backend-policy-allow-ingress-pod-namespace-label.yaml"), nwpolicyName, namespace)
					Expect(err).NotTo(HaveOccurred())

					By("Ensuring we don't have ingress access from role:frontend pods in production namespace")
					pl = pod.List{Pods: frontendProdPods}
					for _, backendPod := range backendPods {
						pass, err = pl.ValidateCurlConnection(backendPod.Status.PodIP, 30*time.Second, validateNetworkPolicyTimeout)
						if err != nil {
							e := backendPod.Describe()
							if e != nil {
								log.Printf("Unable to describe pod %s\n: %s", backendPod.Metadata.Name, e)
							}
						}
						Expect(err).Should(HaveOccurred())
						Expect(pass).To(BeFalse())
					}

					By("Ensuring we have ingress access from role:frontend pods in development namespace")
					pl = pod.List{Pods: frontendDevPods}
					for _, backendPod := range backendPods {
						pass, err = pl.ValidateCurlConnection(backendPod.Status.PodIP, 30*time.Second, cfg.Timeout)
						if err != nil {
							e := backendPod.Describe()
							if e != nil {
								log.Printf("Unable to describe pod %s\n: %s", backendPod.Metadata.Name, e)
							}
						}
						Expect(err).NotTo(HaveOccurred())
						Expect(pass).To(BeTrue())
					}

					By("Cleaning up after ourselves")
					networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)
				}

				By("Cleaning up after ourselves")
				err = frontendProdDeployment.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = frontendDevDeployment.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = backendDeployment.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = nwpolicyDeployment.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = namespaceDev.Delete()
				Expect(err).NotTo(HaveOccurred())
				err = namespaceProd.Delete()
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("Calico or Azure network policy was not provisioned for this Cluster Definition")
			}
		})
	})

	Describe("with a windows agent pool", func() {
		It("should be able to deploy and scale an iis webserver", func() {
			if eng.HasWindowsAgents() {
				windowsImages, err := eng.GetWindowsTestImages()
				Expect(err).NotTo(HaveOccurred())
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				deploymentPrefix := fmt.Sprintf("iis-%s", cfg.Name)
				deploymentName := fmt.Sprintf("%s-%v", deploymentPrefix, r.Intn(99999))
				By("Creating a deployment with 1 pod running IIS")
				iisDeploy, err := deployment.CreateWindowsDeployWithHostportDeleteIfExist(deploymentPrefix, windowsImages.IIS, deploymentName, "default", 80, -1)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting on pod to be Ready")
				running, err := pod.WaitOnSuccesses(deploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Exposing a LoadBalancer for the pod")
				err = iisDeploy.ExposeDeleteIfExist(deploymentPrefix, "default", "LoadBalancer", 80, 80)
				Expect(err).NotTo(HaveOccurred())
				iisService, err := service.Get(deploymentName, "default")
				Expect(err).NotTo(HaveOccurred())
				err = iisService.WaitForIngress(cfg.Timeout, 5*time.Second)
				Expect(err).NotTo(HaveOccurred())

				By("Verifying that the service is reachable and returns the default IIS start page")
				err = iisService.ValidateWithRetry("(IIS Windows Server)", sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Checking that each pod can reach the internet")
				var iisPods []pod.Pod
				iisPods, err = iisDeploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(iisPods)).ToNot(BeZero())
				for _, iisPod := range iisPods {
					var pass bool
					pass, err = iisPod.CheckWindowsOutboundConnection(sleepBetweenRetriesWhenWaitingForPodReady, timeoutWhenWaitingForPodOutboundAccess)
					Expect(err).NotTo(HaveOccurred())
					Expect(pass).To(BeTrue())
				}

				By("Scaling deployment to 5 pods")
				err = iisDeploy.ScaleDeployment(5)
				Expect(err).NotTo(HaveOccurred())
				_, err = iisDeploy.WaitForReplicas(5, 5, 2*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting on 5 pods to be Ready")
				running, err = pod.WaitOnSuccesses(deploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				iisPods, err = iisDeploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(iisPods)).To(Equal(5))

				By("Verifying that the service is reachable and returns the default IIS start page")
				err = iisService.ValidateWithRetry("(IIS Windows Server)", sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Checking that each pod can reach the internet")
				iisPods, err = iisDeploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(iisPods)).ToNot(BeZero())
				for _, iisPod := range iisPods {
					var pass bool
					pass, err = iisPod.CheckWindowsOutboundConnection(sleepBetweenRetriesWhenWaitingForPodReady, timeoutWhenWaitingForPodOutboundAccess)
					Expect(err).NotTo(HaveOccurred())
					Expect(pass).To(BeTrue())
				}

				By("Scaling deployment to 2 pods")
				err = iisDeploy.ScaleDeployment(2)
				Expect(err).NotTo(HaveOccurred())
				_, err = iisDeploy.WaitForReplicas(2, 2, 2*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				iisPods, err = iisDeploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(iisPods)).To(Equal(2))

				By("Verifying that the service is reachable and returns the default IIS start page")
				err = iisService.ValidateWithRetry("(IIS Windows Server)", sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Checking that each pod can reach the internet")
				iisPods, err = iisDeploy.PodsRunning()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(iisPods)).ToNot(BeZero())
				for _, iisPod := range iisPods {
					var pass bool
					pass, err = iisPod.CheckWindowsOutboundConnection(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(pass).To(BeTrue())
				}

				By("Verifying pods & services can be deleted")
				err = iisDeploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = iisService.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("No windows agent was provisioned for this Cluster Definition")
			}
		})

		It("should be able to resolve DNS across windows and linux deployments", func() {
			if eng.HasWindowsAgents() {
				windowsImages, err := eng.GetWindowsTestImages()
				Expect(err).NotTo(HaveOccurred())
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				deploymentPrefix := fmt.Sprintf("iis-dns-%s", cfg.Name)
				windowsDeploymentName := fmt.Sprintf("%s-%v", deploymentPrefix, r.Intn(99999))
				By("Creating a deployment running IIS")
				windowsIISDeployment, err := deployment.CreateWindowsDeployWithHostportDeleteIfExist(deploymentPrefix, windowsImages.IIS, windowsDeploymentName, "default", 80, -1)
				Expect(err).NotTo(HaveOccurred())

				deploymentPrefix = fmt.Sprintf("nginx-dns-%s", cfg.Name)
				nginxDeploymentName := fmt.Sprintf("%s-%v", deploymentPrefix, r.Intn(99999))
				By("Creating a nginx deployment")
				linuxNginxDeploy, err := deployment.CreateLinuxDeployDeleteIfExists(deploymentPrefix, "library/nginx:latest", nginxDeploymentName, "default", "")
				Expect(err).NotTo(HaveOccurred())

				By("Ensure there is a Running nginx pod")
				running, err := pod.WaitOnSuccesses(nginxDeploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Ensure there is a Running iis pod")
				running, err = pod.WaitOnSuccesses(windowsDeploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))

				By("Exposing a internal service for the linux nginx deployment")
				err = linuxNginxDeploy.ExposeIfNotExist("ClusterIP", 80, 80)
				Expect(err).NotTo(HaveOccurred())
				linuxService, err := service.Get(nginxDeploymentName, "default")
				Expect(err).NotTo(HaveOccurred())

				By("Exposing a internal service for the windows iis deployment")
				err = windowsIISDeployment.ExposeIfNotExist("ClusterIP", 80, 80)
				Expect(err).NotTo(HaveOccurred())
				windowsService, err := service.Get(windowsDeploymentName, "default")
				Expect(err).NotTo(HaveOccurred())

				By("Connecting to Windows from another Windows deployment")
				name := fmt.Sprintf("windows-2-windows-%s", cfg.Name)
				command := fmt.Sprintf("iwr -UseBasicParsing -TimeoutSec 60 %s", windowsService.Metadata.Name)
				successes, err := pod.RunCommandMultipleTimes(pod.RunWindowsPod, windowsImages.ServerCore, name, command, cfg.StabilityIterations, 1*time.Second, windowsCommandTimeout, retryCommandsTimeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))

				By("Connecting to Linux from Windows deployment")
				name = fmt.Sprintf("windows-2-linux-%s", cfg.Name)
				command = fmt.Sprintf("iwr -UseBasicParsing -TimeoutSec 60 %s", linuxService.Metadata.Name)
				successes, err = pod.RunCommandMultipleTimes(pod.RunWindowsPod, windowsImages.ServerCore, name, command, cfg.StabilityIterations, 1*time.Second, windowsCommandTimeout, retryCommandsTimeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))

				By("Connecting to Windows from Linux deployment")
				name = fmt.Sprintf("linux-2-windows-%s", cfg.Name)
				command = fmt.Sprintf("wget %s", windowsService.Metadata.Name)
				successes, err = pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, windowsCommandTimeout, retryCommandsTimeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))

				By("Cleaning up after ourselves")
				err = windowsIISDeployment.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = linuxNginxDeploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = windowsService.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = linuxService.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("No windows agent was provisioned for this Cluster Definition")
			}
		})

		// Windows Bug 18213017: Kubernetes Hostport mappings don't work
		/*
			It("should be able to reach hostport in an iis webserver", func() {
				if eng.HasWindowsAgents() {
					r := rand.New(rand.NewSource(time.Now().UnixNano()))
					hostport := 8123
					deploymentName := fmt.Sprintf("iis-%s-%v", cfg.Name, r.Intn(99999))
					iisDeploy, err := deployment.CreateWindowsDeployIfNotExist(iisImage, deploymentName, "default", 80, hostport)
					Expect(err).NotTo(HaveOccurred())
					running, err := pod.WaitOnSuccesses(deploymentName, "default", 4, 30*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(running).To(Equal(true))
					iisPods, err := iisDeploy.PodsRunning()
					Expect(err).NotTo(HaveOccurred())
					Expect(len(iisPods)).ToNot(BeZero())
					kubeConfig, err := GetConfigWithRetry(3*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					master := fmt.Sprintf("azureuser@%s", kubeConfig.GetServerName())
					for _, iisPod := range iisPods {
						valid := iisPod.ValidateHostPort("(IIS Windows Server)", 10, 10*time.Second, master, masterSSHPrivateKeyFilepath)
						Expect(valid).To(BeTrue())
					}
					err = iisDeploy.Delete(kubectlOutput)
					Expect(err).NotTo(HaveOccurred())
				} else {
					Skip("No windows agent was provisioned for this Cluster Definition")
				}
			})*/
		It("should be able to attach azure file", func() {
			if eng.HasWindowsAgents() {
				if eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion == "1.11.0" {
					// Failure in 1.11.0 - https://github.com/kubernetes/kubernetes/issues/65845, fixed in 1.11.1
					Skip("Kubernetes 1.11.0 has a known issue creating Azure PersistentVolumeClaim")
				} else if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.8.0") {
					windowsImages, err := eng.GetWindowsTestImages()
					Expect(err).NotTo(HaveOccurred())

					iisAzurefileYaml, err := pod.ReplaceContainerImageFromFile(filepath.Join(WorkloadDir, "iis-azurefile.yaml"), windowsImages.IIS)
					Expect(err).NotTo(HaveOccurred())
					defer os.Remove(iisAzurefileYaml)

					By("Creating an AzureFile storage class")
					storageclassName := "azurefile" // should be the same as in storageclass-azurefile.yaml
					sc, err := storageclass.CreateStorageClassFromFile(filepath.Join(WorkloadDir, "storageclass-azurefile.yaml"), storageclassName)
					Expect(err).NotTo(HaveOccurred())
					ready, err := sc.WaitOnReady(5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))

					By("Creating a persistent volume claim")
					pvcName := "pvc-azurefile" // should be the same as in pvc-azurefile.yaml
					pvc, err := persistentvolumeclaims.CreatePVCFromFileDeleteIfExist(filepath.Join(WorkloadDir, "pvc-azurefile.yaml"), pvcName, "default")
					Expect(err).NotTo(HaveOccurred())
					ready, err = pvc.WaitOnReady("default", 5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))

					By("Launching an IIS pod using the volume claim")
					podName := "iis-azurefile" // should be the same as in iis-azurefile.yaml
					iisPod, err := pod.CreatePodFromFile(iisAzurefileYaml, podName, "default", 1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					ready, err = iisPod.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))

					By("Checking that the pod can access volume")
					valid, err := iisPod.ValidateAzureFile("mnt\\azure", 10*time.Second, 3*time.Minute)
					Expect(valid).To(BeTrue())
					Expect(err).NotTo(HaveOccurred())

					err = iisPod.Delete(util.DefaultDeleteRetries)
					Expect(err).NotTo(HaveOccurred())
					err = pvc.Delete(util.DefaultDeleteRetries)
					Expect(err).NotTo(HaveOccurred())
				} else {
					Skip("Kubernetes version needs to be 1.8 and up for Azure File test")
				}
			} else {
				Skip("No windows agent was provisioned for this Cluster Definition")
			}
		})
	})

	Describe("after the cluster has been up for awhile", func() {
		It("dns-liveness pod should not have any restarts", func() {
			pod, err := pod.Get("dns-liveness", "default", podLookupRetries)
			Expect(err).NotTo(HaveOccurred())
			running, err := pod.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, 3*time.Minute)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))
			restarts := pod.Status.ContainerStatuses[0].RestartCount
			if cfg.SoakClusterName == "" {
				err = pod.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				Expect(restarts).To(Equal(0))
			} else {
				log.Printf("%d DNS livenessProbe restarts since this cluster was created...\n", restarts)
			}
		})

		It("should have healthy time synchronization", func() {
			if !eng.ExpandedDefinition.Properties.HasLowPriorityScaleset() {
				nodeList, err := node.GetReady()
				Expect(err).NotTo(HaveOccurred())
				timeSyncValidateScript := "time-sync-validate.sh"
				err = sshConn.CopyTo(timeSyncValidateScript)
				Expect(err).NotTo(HaveOccurred())
				timeSyncValidationCommand := fmt.Sprintf("\"/tmp/%s\"", timeSyncValidateScript)
				err = sshConn.Execute(timeSyncValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
				for _, node := range nodeList.Nodes {
					if node.IsUbuntu() && !firstMasterRegexp.MatchString(node.Metadata.Name) {
						err := sshConn.CopyToRemote(node.Metadata.Name, "/tmp/"+timeSyncValidateScript)
						Expect(err).NotTo(HaveOccurred())
						err = sshConn.ExecuteRemoteWithRetry(node.Metadata.Name, timeSyncValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should be able to cleanup the long running php-apache stuff", func() {
			if cfg.SoakClusterName == "" {
				phpApacheDeploy, err := deployment.GetWithRetry(longRunningApacheDeploymentName, "default", 3*time.Second, 1*time.Minute)
				if err != nil {
					fmt.Println(err)
				}
				Expect(err).NotTo(HaveOccurred())
				s, err := service.Get(longRunningApacheDeploymentName, "default")
				Expect(err).NotTo(HaveOccurred())

				err = s.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = phpApacheDeploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("Keep long-running php-apache workloads running for soak clusters")
			}
		})
	})
})
