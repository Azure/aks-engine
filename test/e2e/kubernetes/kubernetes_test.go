//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/engine"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/daemonset"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/deployment"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/event"
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
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	WorkloadDir                               = "workloads"
	PolicyDir                                 = "workloads/policies"
	kubeSystemPodsReadinessChecks             = 6
	sleepBetweenRetriesWhenWaitingForPodReady = 1 * time.Second
	sleepBetweenRetriesRemoteSSHCommand       = 3 * time.Second
	timeoutWhenWaitingForPodOutboundAccess    = 1 * time.Minute
	singleCommandTimeout                      = 1 * time.Minute
	validateNetworkPolicyTimeout              = 3 * time.Minute
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
	masterNodes                     []node.Node
	clusterAutoscalerEngaged        bool
	clusterAutoscalerAddon          api.KubernetesAddon
	deploymentReplicasCount         int
	dnsAddonName                    string
	stabilityCommandTimeout         time.Duration
	env                             azure.Environment
	azureClient                     *armhelpers.AzureClient
	firstMasterRegexStr             = fmt.Sprintf("^%s-.*-0", common.LegacyControlPlaneVMPrefix)
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
	longRunningApacheDeploymentName = "php-apache-long-running"
	for _, profile := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
		deploymentReplicasCount += profile.Count
	}

	var getNodeByRegexError error
	masterNodes, getNodeByRegexError = node.GetByRegexWithRetry(fmt.Sprintf("^%s-", common.LegacyControlPlaneVMPrefix), 3*time.Minute, cfg.Timeout)
	Expect(getNodeByRegexError).NotTo(HaveOccurred())
	var getKubeConfigError error
	kubeConfig, getKubeConfigError = GetConfigWithRetry(3*time.Second, cfg.Timeout)
	Expect(getKubeConfigError).NotTo(HaveOccurred())

	if cfg.RebootControlPlaneNodes {
		cfg.BlockSSHPort = true
		cfg.StabilityIterations = 0
	}

	if !cfg.BlockSSHPort {
		var err error
		masterName := masterNodes[0].Metadata.Name
		if strings.Contains(masterName, "vmss") {
			masterSSHPort = "50001"
		} else {
			masterSSHPort = "22"
		}
		masterSSHPrivateKeyFilepath = cfg.GetSSHKeyPath()
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
	}
	if hasAddon, addon := eng.HasAddon(common.ClusterAutoscalerAddonName); hasAddon {
		clusterAutoscalerAddon = addon
		if len(addon.Pools) > 0 {
			for _, pool := range addon.Pools {
				p := eng.ExpandedDefinition.Properties.GetAgentPoolIndexByName(pool.Name)
				maxNodes, _ := strconv.Atoi(pool.Config["max-nodes"])
				minNodes, _ := strconv.Atoi(pool.Config["min-nodes"])
				if maxNodes > eng.ExpandedDefinition.Properties.AgentPoolProfiles[p].Count &&
					minNodes <= eng.ExpandedDefinition.Properties.AgentPoolProfiles[p].Count {
					clusterAutoscalerEngaged = true
					break
				}
			}
		}
	}
	if hasAddon, _ := eng.HasAddon(common.KubeDNSAddonName); hasAddon {
		dnsAddonName = common.KubeDNSAddonName
	}
	if hasAddon, _ := eng.HasAddon("coredns"); hasAddon {
		dnsAddonName = common.CoreDNSAddonName
	}
	Expect(dnsAddonName).NotTo(Equal(""))

	stabilityCommandTimeout = time.Duration(cfg.StabilityTimeoutSeconds) * time.Second

	if !cfg.IsCustomCloudProfile() {
		env, err = azure.EnvironmentFromName("AzurePublicCloud") // TODO get this programmatically
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
		azureClient, err = armhelpers.NewAzureClientWithClientSecret(env, cfg.SubscriptionID, cfg.ClientID, cfg.ClientSecret)
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
	}
})

var _ = AfterSuite(func() {
	if cfg.DebugAfterSuite {
		cmd := exec.Command("k", "get", "deployments,pods,svc,daemonsets,configmaps,endpoints,jobs,clusterroles,clusterrolebindings,roles,rolebindings,storageclasses,podsecuritypolicy", "--all-namespaces", "-o", "wide")
		out, err := cmd.CombinedOutput()
		log.Printf("%s\n", out)
		if err != nil {
			log.Printf("Error: Unable to print all cluster resources\n")
		}
		pod.PrintPodsLogs("kube-addon-manager", "kube-system", 5*time.Second, 1*time.Minute)
		pod.PrintPodsLogs("kube-proxy", "kube-system", 5*time.Second, 1*time.Minute)
		pod.PrintPodsLogs("kube-scheduler", "kube-system", 5*time.Second, 1*time.Minute)
		pod.PrintPodsLogs(common.APIServerComponentName, "kube-system", 5*time.Second, 1*time.Minute)
		pod.PrintPodsLogs("kube-controller-manager", "kube-system", 5*time.Second, 1*time.Minute)
	}
})

var _ = Describe("Azure Container Cluster using the Kubernetes Orchestrator", func() {
	Describe("regardless of agent pool type", func() {
		It("should check for cluster-init pod", func() {
			if cfg.ClusterInitPodName != "" {
				By(fmt.Sprintf("Ensuring that cluster-init Pod \"%s\" is Running", cfg.ClusterInitPodName))
				running, err := pod.WaitOnSuccesses(cfg.ClusterInitPodName, "default", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
			}
			if cfg.ClusterInitJobName != "" {
				ready, err := job.WaitOnSucceeded(cfg.ClusterInitJobName, "default", 30*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(ready).To(Equal(true))
			}
		})

		It("should validate filesystem config", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				filesystemValidateScript := "host-os-fs.sh"
				err = sshConn.CopyTo(filesystemValidateScript)
				Expect(err).NotTo(HaveOccurred())
				envString := fmt.Sprintf("MASTER_NODE=true")
				filesystemValidationCommand := fmt.Sprintf("%s /tmp/%s", envString, filesystemValidateScript)
				err = sshConn.Execute(filesystemValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
						err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+filesystemValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						envString = fmt.Sprintf("MASTER_NODE=%t", n.HasSubstring([]string{common.LegacyControlPlaneVMPrefix}))
						filesystemValidationCommand = fmt.Sprintf("%s /tmp/%s", envString, filesystemValidateScript)
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, filesystemValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate host OS DNS", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else {
				var nodes []node.Node
				var err error
				if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
					nodes, err = node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				} else {
					nodes = masterNodes
				}
				Expect(err).NotTo(HaveOccurred())
				hostOSDNSValidateScript := "host-os-dns-validate.sh"
				err = sshConn.CopyTo(hostOSDNSValidateScript)
				Expect(err).NotTo(HaveOccurred())
				envString := "NODE_HOSTNAMES='"
				for _, n := range nodes {
					envString += fmt.Sprintf("%s ", n.Metadata.Name)
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
				for _, n := range nodes {
					if n.IsLinux() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
						err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+hostOSDNSValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, hostOSDNSValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			}
		})

		It("should validate cloudprovider config", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				var cloudproviderEnabledPrefixes []string
				if eng.ExpandedDefinition.Properties.MasterProfile != nil {
					cloudproviderEnabledPrefixes = append(cloudproviderEnabledPrefixes, fmt.Sprintf("%s-", common.LegacyControlPlaneVMPrefix))
				}
				for _, profile := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
					if profile.RequiresCloudproviderConfig() {
						cloudproviderEnabledPrefixes = append(cloudproviderEnabledPrefixes, "k8s-"+profile.Name)
					}
				}
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				cloudproviderConfigValidateScript := "cloudprovider-config-validate.sh"
				err = sshConn.CopyTo(cloudproviderConfigValidateScript)
				Expect(err).NotTo(HaveOccurred())
				envString := fmt.Sprintf("BACKOFF_MODE=%s", eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoffMode)
				// TODO test remaining cloudprovider config
				cloudproviderConfigValidationCommand := fmt.Sprintf("%s /tmp/%s", envString, cloudproviderConfigValidateScript)
				err = sshConn.Execute(cloudproviderConfigValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) && n.HasSubstring(cloudproviderEnabledPrefixes) {
						err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+cloudproviderConfigValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, cloudproviderConfigValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should have the expected k8s version", func() {
			customHyperkubeImage := eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage
			customWindowsPackageURL := eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.CustomWindowsPackageURL
			if customHyperkubeImage == "" && customWindowsPackageURL == "" {
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					err := n.Describe()
					if err != nil {
						log.Printf("Unable to describe node %s: %s", n.Metadata.Name, err)
					}
					Expect("v" + eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(n.Version()))
				}
			} else if customHyperkubeImage != "" {
				customHyperkubeValidateScript := "custom-hyperkube-validate.sh"
				err := sshConn.CopyTo(customHyperkubeValidateScript)
				Expect(err).NotTo(HaveOccurred())
				envString := fmt.Sprintf("CUSTOM_HYPERKUBE_IMAGE=%s", customHyperkubeImage)
				customHyperkubeValidationCommand := fmt.Sprintf("%s /tmp/%s", envString, customHyperkubeValidateScript)
				err = sshConn.Execute(customHyperkubeValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("This is a cluster built from source")
			}
		})

		It("should display the installed Ubuntu version on the master node", func() {
			if !eng.ExpandedDefinition.Properties.MasterProfile.IsUbuntu() {
				Skip("This is not an ubuntu master")
			} else if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else {
				lsbReleaseCmd := fmt.Sprintf("lsb_release -a && uname -r")
				err := sshConn.Execute(lsbReleaseCmd, true)
				Expect(err).NotTo(HaveOccurred())
				kernelVerCmd := fmt.Sprintf("cat /proc/version")
				err = sshConn.Execute(kernelVerCmd, true)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("should display the installed docker runtime on all nodes", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.RequiresDocker() {
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					dockerVersionCmd := fmt.Sprintf("\"docker version\"")
					for _, n := range nodes {
						if n.IsWindows() {
							if eng.ExpandedDefinition.Properties.WindowsProfile != nil && !eng.ExpandedDefinition.Properties.WindowsProfile.GetSSHEnabled() {
								log.Printf("Can't ssh into Windows node %s because there is no SSH listener", n.Metadata.Name)
								continue
							}
						}
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, dockerVersionCmd, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
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
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					rootPasswdCmd := fmt.Sprintf("\"sudo grep '^root:[!*]:' /etc/shadow\" && exit 1 || exit 0")
					for _, n := range nodes {
						if n.IsUbuntu() {
							err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, rootPasswdCmd, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
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

		It("should validate aks-engine-provided sysctl configuration", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				for key, val := range eng.ExpandedDefinition.Properties.MasterProfile.SysctlDConfig {
					for _, n := range nodes {
						if n.HasSubstring([]string{common.LegacyControlPlaneVMPrefix}) && n.IsUbuntu() {
							err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, fmt.Sprintf("sysctl %s | grep '= %s'", key, val), false, sleepBetweenRetriesRemoteSSHCommand, singleCommandTimeout)
							Expect(err).NotTo(HaveOccurred())
						}
					}
				}
				for _, pool := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
					for key, val := range pool.SysctlDConfig {
						for _, n := range nodes {
							if n.HasSubstring([]string{pool.Name}) && n.IsUbuntu() {
								err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, fmt.Sprintf("sysctl %s | grep '= %s'", key, val), false, sleepBetweenRetriesRemoteSSHCommand, singleCommandTimeout)
								Expect(err).NotTo(HaveOccurred())
							}
						}
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate Ubuntu host OS network configuration on all nodes", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					var largeSKUPrefixes []string
					if eng.ExpandedDefinition.Properties.MasterProfile != nil {
						if util.IsLargeVMSKU(eng.ExpandedDefinition.Properties.MasterProfile.VMSize) {
							largeSKUPrefixes = append(largeSKUPrefixes, fmt.Sprintf("%s-", common.LegacyControlPlaneVMPrefix))
						}
					}
					for _, profile := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
						if util.IsLargeVMSKU(profile.VMSize) {
							largeSKUPrefixes = append(largeSKUPrefixes, "k8s-"+profile.Name)
						}
					}
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					netConfigValidateScript := "net-config-validate.sh"
					err = sshConn.CopyTo(netConfigValidateScript)
					Expect(err).NotTo(HaveOccurred())
					for _, n := range nodes {
						var gt8CoreSKU string
						if n.HasSubstring(largeSKUPrefixes) && n.IsUbuntu() {
							gt8CoreSKU = "true"
						}
						netConfigValidationCommand := fmt.Sprintf("\"GT_8_CORE_SKU=%s /tmp/%s\"", gt8CoreSKU, netConfigValidateScript)
						if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
							err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+netConfigValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, netConfigValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
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
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					CISFilesValidateScript := "CIS-files-validate.sh"
					err = sshConn.CopyTo(CISFilesValidateScript)
					Expect(err).NotTo(HaveOccurred())
					CISValidationCommand := fmt.Sprintf("\"/tmp/%s\"", CISFilesValidateScript)
					err = sshConn.Execute(CISValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, n := range nodes {
						if !firstMasterRegexp.MatchString(n.Metadata.Name) {
							err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+CISFilesValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, CISValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
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
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					modprobeConfigValidateScript := "modprobe-config-validate.sh"
					err = sshConn.CopyTo(modprobeConfigValidateScript)
					Expect(err).NotTo(HaveOccurred())
					netConfigValidationCommand := fmt.Sprintf("\"/tmp/%s\"", modprobeConfigValidateScript)
					err = sshConn.Execute(netConfigValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, n := range nodes {
						if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
							err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+modprobeConfigValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, netConfigValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
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
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				installedPackagesValidateScript := "ubuntu-installed-packages-validate.sh"
				err = sshConn.CopyTo(installedPackagesValidateScript)
				Expect(err).NotTo(HaveOccurred())
				installedPackagesValidationCommand := fmt.Sprintf("\"/tmp/%s\"", installedPackagesValidateScript)
				err = sshConn.Execute(installedPackagesValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
						err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+installedPackagesValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, installedPackagesValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should validate that every linux node has the right sshd config", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					sshdConfigValidateScript := "sshd-config-validate.sh"
					err = sshConn.CopyTo(sshdConfigValidateScript)
					Expect(err).NotTo(HaveOccurred())
					sshdConfigValidationCommand := fmt.Sprintf("\"/tmp/%s\"", sshdConfigValidateScript)
					err = sshConn.Execute(sshdConfigValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, n := range nodes {
						if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
							err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+sshdConfigValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, sshdConfigValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
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
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.ExpandedDefinition.Properties.IsVHDDistroForAllNodes() {
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					pwQualityValidateScript := "pwquality-validate.sh"
					err = sshConn.CopyTo(pwQualityValidateScript)
					Expect(err).NotTo(HaveOccurred())
					pwQualityValidationCommand := fmt.Sprintf("\"/tmp/%s\"", pwQualityValidateScript)
					err = sshConn.Execute(pwQualityValidationCommand, false)
					Expect(err).NotTo(HaveOccurred())
					for _, n := range nodes {
						if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
							err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+pwQualityValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
							Expect(err).NotTo(HaveOccurred())
							err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, pwQualityValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
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
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else {
				var auditDNodePrefixes []string
				var nonRegularPriVMSSPrefixes []string
				if eng.ExpandedDefinition.Properties.MasterProfile != nil {
					if to.Bool(eng.ExpandedDefinition.Properties.MasterProfile.AuditDEnabled) {
						auditDNodePrefixes = append(auditDNodePrefixes, fmt.Sprintf("%s-", common.LegacyControlPlaneVMPrefix))
					}
				}
				for _, profile := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
					if profile.IsLowPriorityScaleSet() || profile.IsSpotScaleSet() {
						nonRegularPriVMSSPrefixes = append(nonRegularPriVMSSPrefixes, "k8s-"+profile.Name)
					} else if to.Bool(profile.AuditDEnabled) {
						auditDNodePrefixes = append(auditDNodePrefixes, profile.Name)
					}
				}
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				auditdValidateScript := "auditd-validate.sh"
				err = sshConn.CopyTo(auditdValidateScript)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					if !n.HasSubstring(nonRegularPriVMSSPrefixes) && n.IsUbuntu() {
						var enabled bool
						if n.HasSubstring(auditDNodePrefixes) {
							enabled = true
						}
						err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+auditdValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						auditdValidationCommand := fmt.Sprintf("\"ENABLED=%t /tmp/%s\"", enabled, auditdValidateScript)
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, auditdValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			}
		})

		It("should report all nodes in a Ready state", func() {
			var expectedReadyNodes int
			if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() &&
				!clusterAutoscalerEngaged &&
				cfg.AddNodePoolInput == "" {
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

		It("should have node labels and annotations added by E2E test runner", func() {
			if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() &&
				cfg.AddNodePoolInput == "" && !cfg.RebootControlPlaneNodes {
				totalNodeCount := eng.NodeCount()
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

		It("should have core kube-system componentry running", func() {
			coreComponents := []string{
				common.AddonManagerComponentName,
				common.APIServerComponentName,
				common.ControllerManagerComponentName,
				common.KubeProxyAddonName,
				common.SchedulerComponentName,
			}
			if to.Bool(eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager) {
				coreComponents = append(coreComponents, common.CloudControllerManagerComponentName)
			}
			if to.Bool(eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms) {
				coreComponents = append(coreComponents, common.AzureKMSProviderComponentName)
			}
			for _, componentName := range coreComponents {
				By(fmt.Sprintf("Ensuring that %s is Running", componentName))
				running, err := pod.WaitOnSuccesses(componentName, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
			}

			customHyperkubeImage := eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage
			if customHyperkubeImage != "" {
				hyperkubeComponents := []string{
					common.APIServerComponentName,
					common.ControllerManagerComponentName,
					common.KubeProxyAddonName,
					common.SchedulerComponentName,
				}

				for _, hyperkubeComponent := range hyperkubeComponents {
					pods, err := pod.GetAllByPrefixWithRetry(hyperkubeComponent, "kube-system", 3*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					for _, pod := range pods {
						Expect(pod.Spec.Containers[0].Image).To(Equal(customHyperkubeImage))
					}
				}
			}
		})

		It("should be able to schedule a pod to a control plane node", func() {
			By("Creating a Job with control plane nodeSelector")
			for i := 1; i <= 3; i++ {
				j, err := job.CreateJobFromFileWithRetry(filepath.Join(WorkloadDir, "busybox-master.yaml"), "busybox-master", "default", 3*time.Second, 3*time.Minute)
				if err != nil {
					fmt.Printf("unable to create job: %s\n", err)
					continue
				}
				ready, err := j.WaitOnSucceeded(30*time.Second, 3*time.Minute)
				if err != nil {
					fmt.Printf("timed out waiting for pod success: %s\n", err)
					continue
				}
				Expect(ready).To(Equal(true))
				fmt.Printf("successfully scheduled a pod to the control plane in %d attempts\n", i)
				break
			}
		})

		It("should be able to schedule a pod to a Linux node", func() {
			if eng.AnyAgentIsLinux() {
				By("Creating a Job with agent nodeSelector")
				for i := 1; i <= 3; i++ {
					j, err := job.CreateJobFromFileWithRetry(filepath.Join(WorkloadDir, "busybox-agent.yaml"), "busybox-agent", "default", 3*time.Second, 3*time.Minute)
					if err != nil {
						fmt.Printf("unable to create job: %s\n", err)
						continue
					}
					ready, err := j.WaitOnSucceeded(30*time.Second, 3*time.Minute)
					if err != nil {
						fmt.Printf("timed out waiting for pod success: %s\n", err)
						continue
					}
					Expect(ready).To(Equal(true))
					fmt.Printf("successfully scheduled a pod to a Linux node in %d attempts\n", i)
					break
				}
			} else {
				Skip("agent nodeSelector test Job is currently Linux-only")
			}
		})

		It("should be able to schedule a pod to a Windows node", func() {
			if eng.HasWindowsAgents() {
				windowsImages, err := eng.GetWindowsTestImages()
				Expect(err).NotTo(HaveOccurred())
				p, err := pod.RunWindowsWithRetry(windowsImages.ServerCore, "windows-schedule-test", "default", "powershell", true, 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				succeeded, err := p.WaitOnSucceeded(10*time.Second, cfg.Timeout)
				Expect(succeeded).To(Equal(true))
				//err = pod.Delete(util.DefaultDeleteRetries)
				//Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("no Windows nodes")
			}
		})

		It("should have core kube-system addons running the correct version", func() {
			if eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.CustomKubeProxyImage == "" {
				By(fmt.Sprintf("Ensuring that the %s addon image matches orchestrator version", common.KubeProxyAddonName))
				ds, err := daemonset.Get(common.KubeProxyAddonName, "kube-system", 3)
				Expect(err).NotTo(HaveOccurred())
				log.Printf("Image: %s", ds.Spec.Template.TemplateSpec.Containers[0].Image)
				log.Printf("OrchestratorVersion: %s", eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion)
				version := eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion
				Expect(strings.Contains(ds.Spec.Template.TemplateSpec.Containers[0].Image, version)).To(Equal(true))
			} else {
				Skip("Skipping as testing custom kube-proxy image")
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
					if strings.Contains(currentPod.Metadata.Name, common.ClusterAutoscalerAddonName) {
						log.Print("need to investigate cluster-autoscaler restarts!")
						tooManyRestarts = 10
					}
					Expect(currentPod.Status.ContainerStatuses[0].RestartCount).To(BeNumerically("<", tooManyRestarts))
				}
			} else {
				Skip("Skipping this DEBUG test")
			}
		})

		It("should print cluster resources", func() {
			cmd := exec.Command("k", "get", "deployments,pods,svc,daemonsets,configmaps,endpoints,jobs,clusterroles,clusterrolebindings,roles,rolebindings,storageclasses,podsecuritypolicy", "--all-namespaces", "-o", "wide")
			out, err := cmd.CombinedOutput()
			log.Printf("%s\n", out)
			if err != nil {
				log.Printf("Error: Unable to print all cluster resources\n")
			}
		})

		It("should have DNS resolver pod running", func() {
			By(fmt.Sprintf("Ensuring that %s is running", dnsAddonName))
			running, err := pod.WaitOnSuccesses(dnsAddonName, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			pod.PrintPodsLogs(dnsAddonName, "kube-system", 5*time.Second, 1*time.Minute)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))
		})

		It("should have functional container networking DNS", func() {
			By("Ensuring that we have functional DNS resolution from a linux container")
			validateDNSLinuxName := "validate-dns-linux"
			validateDNSLinuxNamespace := "default"
			j, err := job.CreateJobFromFileWithRetry(filepath.Join(WorkloadDir, fmt.Sprintf("%s.yaml", validateDNSLinuxName)), validateDNSLinuxName, validateDNSLinuxNamespace, 3*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			ready, err := j.WaitOnSucceeded(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			if err != nil {
				pod.PrintPodsLogs(validateDNSLinuxName, validateDNSLinuxNamespace, 5*time.Second, 1*time.Minute)
				pods, err := pod.GetAllByPrefixWithRetry(validateDNSLinuxName, validateDNSLinuxNamespace, 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				for _, p := range pods {
					out, err := p.Exec("--", "cat", "/etc/resolv.conf")
					log.Printf("%s\n", string(out))
					Expect(err).NotTo(HaveOccurred())
					out, err = p.Exec("--", "ifconfig")
					log.Printf("%s\n", string(out))
					Expect(err).NotTo(HaveOccurred())
					out, err = p.Exec("--", "nc", "-vz", eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP, "53")
					log.Printf("%s\n", string(out))
					Expect(err).NotTo(HaveOccurred())
				}
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
				if err != nil {
					pod.PrintPodsLogs("validate-dns-windows", "default", 5*time.Second, 1*time.Minute)
				}
				Expect(err).NotTo(HaveOccurred())
				Expect(ready).To(Equal(true))
			}

			By("Ensuring that we have stable and responsive DNS resolution")
			p, err := pod.CreatePodFromFileIfNotExist(filepath.Join(WorkloadDir, "dns-loop.yaml"), "dns-loop", "default", 1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			running, err := p.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))
			out, err := p.Exec("--", "dns-loop", "100", "google.com", "0.3", "5")
			log.Printf("%s\n", string(out))
			Expect(err).NotTo(HaveOccurred())

			By("Ensuring that we have stable external DNS resolution as we recycle a bunch of pods")
			name := fmt.Sprintf("alpine-%s", cfg.Name)
			command := fmt.Sprintf("time nc -vz bbc.co.uk 80 || nc -vz google.com 443 || nc -vz microsoft.com 80")
			deploymentCommand := fmt.Sprintf("%s && while true; do sleep 1; done || echo unable to make external connections or resolve dns", command)
			// Ensure across all nodes
			successes, err := deployment.RunDeploymentMultipleTimes(deployment.RunLinuxDeploy, "alpine", name, deploymentCommand, deploymentReplicasCount, cfg.StabilityIterations, 1*time.Second, timeoutWhenWaitingForPodOutboundAccess, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
			successes, err = pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
		})

		It("should be able to launch a long-running container networking DNS liveness pod", func() {
			p, err := pod.CreatePodFromFileIfNotExist(filepath.Join(WorkloadDir, "dns-liveness.yaml"), "dns-liveness", "default", 1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			running, err := p.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))
		})

		It("should be able to run a node reboot daemonset", func() {
			if cfg.RebootControlPlaneNodes {
				_, err := daemonset.CreateDaemonsetFromFileWithRetry(filepath.Join(WorkloadDir, "reboot-control-plane-node.yaml"), "reboot-test", "default", 5*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				pods, err := pod.GetAllRunningByLabelWithRetry("app", "reboot-test", "default", 5*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(pods).NotTo(BeEmpty())
			}
		})

		It("should be able to launch a long running HTTP listener and svc endpoint", func() {
			By("Creating a php-apache deployment")
			phpApacheDeploy, err := deployment.CreateLinuxDeployIfNotExist("deis/hpa-example", longRunningApacheDeploymentName, "default", "", "", 3*time.Second, cfg.Timeout)
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
			// Test for basic UDP networking
			name := fmt.Sprintf("alpine-%s", cfg.Name)
			command := fmt.Sprintf("time nc -vz 8.8.8.8 53 || nc -vz 8.8.4.4 53")
			deploymentCommand := fmt.Sprintf("%s && while true; do sleep 1; done || echo unable to connect externally against known listeners", command)
			// Ensure across all nodes
			successes, err := deployment.RunDeploymentMultipleTimes(deployment.RunLinuxDeploy, "alpine", name, deploymentCommand, deploymentReplicasCount, cfg.StabilityIterations, 1*time.Second, timeoutWhenWaitingForPodOutboundAccess, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
			// Ensure responsiveness
			successes, err = pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))

			// Use curl to test responsive DNS lookup + TCP 443 connectivity
			name = fmt.Sprintf("alpine-%s", cfg.Name)
			command = fmt.Sprintf("time curl --head https://www.bing.com 1> /dev/null || curl --head https://google.com 1> /dev/null || curl --head https://microsoft.com 1> /dev/null")
			deploymentCommand = fmt.Sprintf("%s && while true; do sleep 1; done || echo unable to curl externally against known endpoints", command)
			// Ensure across all nodes
			successes, err = deployment.RunDeploymentMultipleTimes(deployment.RunLinuxDeploy, "byrnedo/alpine-curl", name, deploymentCommand, deploymentReplicasCount, cfg.StabilityIterations, 1*time.Second, timeoutWhenWaitingForPodOutboundAccess, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
			// Ensure responsiveness
			successes, err = pod.RunCommandMultipleTimes(pod.RunLinuxPod, "byrnedo/alpine-curl", name, command, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
		})

		It("should have stable internal container networking as we recycle a bunch of pods", func() {
			name := fmt.Sprintf("alpine-%s", cfg.Name)
			command := fmt.Sprintf("time nc -vz kubernetes 443 && nc -vz kubernetes.default.svc 443 && nc -vz kubernetes.default.svc.cluster.local 443")
			deploymentCommand := fmt.Sprintf("time %s && while true; do sleep 1; done || echo unable to reach internal kubernetes endpoints", command)
			// Ensure across all nodes
			successes, err := deployment.RunDeploymentMultipleTimes(deployment.RunLinuxDeploy, "alpine", name, deploymentCommand, deploymentReplicasCount, cfg.StabilityIterations, 1*time.Second, timeoutWhenWaitingForPodOutboundAccess, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
			// Ensure responsiveness
			successes, err = pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(successes).To(Equal(cfg.StabilityIterations))
		})

		It("should have stable pod-to-pod networking", func() {
			if eng.AnyAgentIsLinux() {
				By("Creating a php-apache deployment")
				phpApacheDeploy, err := deployment.CreateLinuxDeployIfNotExist("deis/hpa-example", longRunningApacheDeploymentName, "default", "", "", 3*time.Second, cfg.Timeout)
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
				By("Creating another pod that will connect to the php-apache pod")
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				commandString := fmt.Sprintf("time nc -vz %s.default.svc.cluster.local 80", longRunningApacheDeploymentName)
				consumerPodName := fmt.Sprintf("consumer-pod-%s-%v", cfg.Name, r.Intn(99999))
				deploymentCommand := fmt.Sprintf("%s && while true; do sleep 1; done || echo unable to connect to in-cluster web listener", commandString)
				// Ensure across all nodes
				successes, err := deployment.RunDeploymentMultipleTimes(deployment.RunLinuxDeploy, "busybox", consumerPodName, deploymentCommand, deploymentReplicasCount, cfg.StabilityIterations, 1*time.Second, timeoutWhenWaitingForPodOutboundAccess, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))
				// Ensure responsiveness
				successes, err = pod.RunCommandMultipleTimes(pod.RunLinuxPod, "busybox", consumerPodName, commandString, cfg.StabilityIterations, 1*time.Second, stabilityCommandTimeout, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))
			} else {
				Skip("Pod-to-pod network tests only valid on Linux clusters")
			}
		})

		It("should have addon pods running", func() {
			timeout := cfg.Timeout
			for _, addonName := range []string{common.CoreDNSAddonName, common.TillerAddonName, common.AADPodIdentityAddonName, common.ACIConnectorAddonName,
				common.AzureDiskCSIDriverAddonName, common.AzureFileCSIDriverAddonName, common.CloudNodeManagerAddonName, common.ClusterAutoscalerAddonName,
				common.BlobfuseFlexVolumeAddonName, common.SMBFlexVolumeAddonName, common.KeyVaultFlexVolumeAddonName, common.DashboardAddonName,
				common.ReschedulerAddonName, common.MetricsServerAddonName, common.NVIDIADevicePluginAddonName, common.ContainerMonitoringAddonName,
				common.AzureCNINetworkMonitorAddonName, common.CalicoAddonName, common.AzureNetworkPolicyAddonName, common.IPMASQAgentAddonName,
				common.AzurePolicyAddonName, common.NodeProblemDetectorAddonName, common.AntreaAddonName, common.FlannelAddonName,
				common.ScheduledMaintenanceAddonName, common.SecretsStoreCSIDriverAddonName} {
				var addonPods = []string{addonName}
				var addonNamespace = "kube-system"
				switch addonName {
				case common.BlobfuseFlexVolumeAddonName:
					addonPods = []string{"blobfuse-flexvol-installer"}
				case common.SMBFlexVolumeAddonName:
					addonPods = []string{"smb-flexvol-installer"}
				case common.ContainerMonitoringAddonName:
					addonPods = []string{"omsagent", "omsagent-rs"}
					if eng.HasWindowsAgents() {
						addonPods = append(addonPods, "omsagent-win")
					}
					timeout = 60 * time.Minute
				case common.AzureNetworkPolicyAddonName:
					addonPods = []string{"azure-npm"}
				case common.DashboardAddonName:
					addonPods = []string{common.DashboardAddonName, "dashboard-metrics-scraper"}
					addonNamespace = common.DashboardAddonName
				case common.AADPodIdentityAddonName:
					addonPods = []string{"nmi", "mic"}
				case common.AzureDiskCSIDriverAddonName:
					addonPods = []string{"csi-azuredisk-node", "csi-azuredisk-controller"}
					if eng.HasWindowsAgents() && common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.18.0") {
						addonPods = append(addonPods, "csi-azuredisk-node-windows")
					}
					if eng.AnyAgentIsLinux() && common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
						addonPods = append(addonPods, "csi-snapshot-controller")
					}
				case common.AzureFileCSIDriverAddonName:
					addonPods = []string{"csi-azurefile-node", "csi-azurefile-controller"}
					if eng.HasWindowsAgents() && common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.18.0") {
						addonPods = append(addonPods, "csi-azurefile-node-windows")
					}
				case common.CloudNodeManagerAddonName:
					addonPods = []string{common.CloudNodeManagerAddonName}
					if eng.HasWindowsAgents() {
						addonPods = append(addonPods, common.CloudNodeManagerAddonName+"-windows")
					}
				case common.CoreDNSAddonName:
					addonPods = []string{common.CoreDNSAddonName, common.CoreDNSAddonName + "-autoscaler"}
				case common.IPMASQAgentAddonName:
					addonPods = []string{"azure-ip-masq-agent"}
				case common.CalicoAddonName:
					addonPods = []string{"calico-node", "calico-typha", "calico-typha-horizontal-autoscaler"}
				case common.AzurePolicyAddonName:
					addonPods = []string{common.AzurePolicyAddonName, "gatekeeper-controller-manager"}
				case common.AntreaAddonName:
					addonPods = []string{common.AntreaAddonName + "-agent", common.AntreaAddonName + "-controller"}
				case common.FlannelAddonName:
					addonPods = []string{"kube-flannel-ds"}
				case common.ScheduledMaintenanceAddonName:
					addonPods = []string{"drainsafe-controller-manager", "drainsafe-controller-scheduledevent-manager"}
				case common.SecretsStoreCSIDriverAddonName:
					addonPods = []string{"csi-secrets-store", "csi-secrets-store-provider-azure"}
				}
				if hasAddon, addon := eng.HasAddon(addonName); hasAddon {
					for _, addonPod := range addonPods {
						if addon.Name == common.AzurePolicyAddonName {
							switch addonPod {
							case common.AzurePolicyAddonName:
								addonNamespace = "kube-system"
							case "gatekeeper-controller-manager":
								addonNamespace = "gatekeeper-system"
							}
						}
						By(fmt.Sprintf("Ensuring that the %s pod(s) in the %s addon is Running", addonPod, addonName))
						running, err := pod.WaitOnSuccesses(addonPod, addonNamespace, kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, timeout)
						Expect(err).NotTo(HaveOccurred())
						Expect(running).To(Equal(true))
					}
				} else {
					fmt.Printf("%s disabled for this cluster, will not test\n", addonName)
				}
			}
		})

		It("should have a working node-problem-detector configuration", func() {
			if hasNpd, _ := eng.HasAddon(common.NodeProblemDetectorAddonName); hasNpd {
				running, err := pod.WaitOnSuccesses(common.NodeProblemDetectorAddonName, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				pods, err := pod.GetAllRunningByPrefixWithRetry(common.NodeProblemDetectorAddonName, "kube-system", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(pods).NotTo(BeEmpty())
				nodeName := pods[0].Spec.NodeName
				// Create a fake kernel message on a node running node-problem-detector
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				msgId := r.Intn(999999999999)
				msg := fmt.Sprintf("kernel: BUG: unable to handle kernel NULL pointer dereference at TESTING-%d", msgId)
				kernelMsgTestCommand := fmt.Sprintf("sudo 'echo %s | sudo tee /dev/kmsg'", msg)
				if cfg.BlockSSHPort {
					Skip("SSH port is blocked")
				}
				err = sshConn.ExecuteRemoteWithRetry(nodeName, kernelMsgTestCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				evt, err := event.GetWithRetry(msg, 5*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(evt.Type).To(Equal("Warning"))
				Expect(evt.Reason).To(Equal("KernelOops"))
			}
		})

		It("should have the correct tiller configuration", func() {
			if hasTiller, tillerAddon := eng.HasAddon(common.TillerAddonName); hasTiller {
				running, err := pod.WaitOnSuccesses(common.TillerAddonName, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
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
			if hasContainerMonitoring, _ := eng.HasAddon(common.ContainerMonitoringAddonName); hasContainerMonitoring {
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

		It("should be able to access the dashboard", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else {
				if hasDashboard, _ := eng.HasAddon(common.DashboardAddonName); hasDashboard {
					By("Ensuring that the kubernetes-dashboard service is Running")
					s, err := service.Get(common.DashboardAddonName, common.DashboardAddonName)
					Expect(err).NotTo(HaveOccurred())
					Expect(s).NotTo(BeNil())
					By("Ensuring that the dashboard responds to requests")
					// start `kubectl proxy` in the background on a random port
					var proxyStdout io.ReadCloser
					var proxyStdoutReader *bufio.Reader
					proxyCmd := exec.Command("k", "proxy", "-p", "0")
					proxyCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
					proxyStdout, err = proxyCmd.StdoutPipe()
					Expect(err).NotTo(HaveOccurred())
					util.PrintCommand(proxyCmd)
					err = proxyCmd.Start()
					Expect(err).NotTo(HaveOccurred())
					defer func() {
						syscall.Kill(-proxyCmd.Process.Pid, syscall.SIGKILL)
						if _, waiterr := proxyCmd.Process.Wait(); waiterr != nil {
							log.Printf("kubectl proxy - wait returned err: %v\n", waiterr)
						}
					}()
					proxyStdoutReader = bufio.NewReader(proxyStdout)
					proxyOutStr, outErr := proxyStdoutReader.ReadString('\n')
					Expect(outErr).NotTo(HaveOccurred())
					log.Printf("kubectl proxy stdout: %s\n", proxyOutStr)
					serverStartPrefix := "Starting to serve on "
					Expect(proxyOutStr).To(HavePrefix(serverStartPrefix))
					dashboardHost := strings.TrimSpace(strings.TrimPrefix(proxyOutStr, serverStartPrefix))
					// get an HTTP response from the dashboard login URL
					url := fmt.Sprintf("http://%s/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login", dashboardHost)
					cmd := exec.Command("curl", "--max-time", "60", "--retry", "10", "--retry-delay", "10", "--retry-max-time", "120", url)
					util.PrintCommand(cmd)
					var out []byte
					out, err = cmd.CombinedOutput()
					log.Printf("%s\n", out)
					Expect(err).NotTo(HaveOccurred())
					Expect(out).To(ContainSubstring("<!doctype html>"))
					Expect(out).To(ContainSubstring("<title>Kubernetes Dashboard</title>"))
				} else {
					Skip("kubernetes-dashboard disabled for this cluster, will not test")
				}
			}
		})

		It("should have the correct storage classes deployed", func() {
			if util.IsUsingEphemeralDisks(eng.ExpandedDefinition.Properties.AgentPoolProfiles) {
				Skip("no storage class is deployed when ephemeral disk is used, will not test")
			}
			var (
				isUsingAzureDiskCSIDriver bool
				isUsingAzureFileCSIDriver bool
				azureDiskProvisioner      string
				azureFileProvisioner      string
			)

			if isUsingAzureDiskCSIDriver, _ = eng.HasAddon(common.AzureDiskCSIDriverAddonName); isUsingAzureDiskCSIDriver {
				azureDiskProvisioner = "disk.csi.azure.com"
			} else {
				azureDiskProvisioner = "kubernetes.io/azure-disk"
			}

			if isUsingAzureFileCSIDriver, _ = eng.HasAddon(common.AzureFileCSIDriverAddonName); isUsingAzureFileCSIDriver {
				azureFileProvisioner = "file.csi.azure.com"
			} else {
				azureFileProvisioner = "kubernetes.io/azure-file"
			}

			azureDiskStorageClasses := []string{"default"}
			// Managed disk is used by default when useCloudControllerManager is enabled
			if to.Bool(eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager) || util.IsUsingManagedDisks(eng.ExpandedDefinition.Properties.AgentPoolProfiles) {
				azureDiskStorageClasses = append(azureDiskStorageClasses, "managed-premium", "managed-standard")
			} else {
				azureDiskStorageClasses = append(azureDiskStorageClasses, "unmanaged-premium", "unmanaged-standard")
			}
			for _, azureDiskStorageClass := range azureDiskStorageClasses {
				sc, err := storageclass.Get(azureDiskStorageClass)
				Expect(err).NotTo(HaveOccurred())
				Expect(sc.Provisioner).To(Equal(azureDiskProvisioner))
				if isUsingAzureDiskCSIDriver && eng.ExpandedDefinition.Properties.HasAvailabilityZones() {
					Expect(sc.VolumeBindingMode).To(Equal("WaitForFirstConsumer"))
					Expect(len(sc.AllowedTopologies)).To(Equal(1))
					Expect(len(sc.AllowedTopologies[0].MatchLabelExpressions)).To(Equal(1))
					Expect(sc.AllowedTopologies[0].MatchLabelExpressions[0].Key).To(Equal("topology.disk.csi.azure.com/zone"))
					for _, zone := range eng.ExpandedDefinition.Properties.AgentPoolProfiles[0].AvailabilityZones {
						Expect(sc.AllowedTopologies[0].MatchLabelExpressions[0].Values).To(ContainElement(eng.ExpandedDefinition.Location + "-" + zone))
					}
				} else {
					Expect(sc.VolumeBindingMode).To(Equal("Immediate"))
				}
				if isUsingAzureDiskCSIDriver && common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
					Expect(sc.AllowVolumeExpansion).To(BeTrue())
				}
			}

			for _, azureFileStorageClass := range []string{"azurefile"} {
				sc, err := storageclass.Get(azureFileStorageClass)
				Expect(err).NotTo(HaveOccurred())
				Expect(sc.Provisioner).To(Equal(azureFileProvisioner))
				Expect(sc.VolumeBindingMode).To(Equal("Immediate"))
				if isUsingAzureFileCSIDriver && common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
					Expect(sc.AllowVolumeExpansion).To(BeTrue())
				}
			}
		})

		It("should be able to kubectl port-forward to a running pod", func() {
			deploymentNamespace := "default"
			testPortForward := func(deploymentName string) {
				running, podWaitErr := pod.WaitOnSuccesses(deploymentName, deploymentNamespace, 3, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(podWaitErr).NotTo(HaveOccurred())
				Expect(running).To(Equal(true))
				d, err := deployment.GetWithRetry(deploymentName, deploymentNamespace, 5*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				pods, err := d.PodsRunning()
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
				deploy, err := deployment.CreateLinuxDeployDeleteIfExists(deploymentPrefix, "library/nginx:latest", deploymentName, deploymentNamespace, "", "", cfg.Timeout)
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
					deploy, err := deployment.CreateWindowsDeployDeleteIfExist(deploymentPrefix, windowsImages.IIS, deploymentName, deploymentNamespace, "", "", cfg.Timeout)
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

		It("should have the correct pods and containers deployed for CSI drivers", func() {
			addons := map[string]string{
				common.AzureDiskCSIDriverAddonName: "azuredisk",
				common.AzureFileCSIDriverAddonName: "azurefile",
			}
			for addonName, shortenedAddonName := range addons {
				if hasAddon, _ := eng.HasAddon(addonName); !hasAddon {
					continue
				}

				// Validate CSI controller pod
				addonPod := fmt.Sprintf("csi-%s-controller", shortenedAddonName)
				containers := []string{"csi-provisioner", "csi-attacher", "liveness-probe", shortenedAddonName}
				if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
					containers = append(containers, "csi-resizer")
				}
				if eng.AnyAgentIsLinux() {
					switch addonName {
					case common.AzureDiskCSIDriverAddonName:
						if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
							containers = append(containers, "csi-snapshotter")
						}
					case common.AzureFileCSIDriverAddonName:
						if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.13.0") &&
							!common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
							containers = append(containers, "csi-snapshotter")
						}
					}
				}
				By(fmt.Sprintf("Ensuring that %s are running within %s pod", containers, addonPod))
				Expect(pod.EnsureContainersRunningInAllPods(containers, addonPod, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)).NotTo(HaveOccurred())

				// Validate CSI node pod
				addonPod = fmt.Sprintf("csi-%s-node", shortenedAddonName)
				containers = []string{"liveness-probe", "node-driver-registrar", shortenedAddonName}
				By(fmt.Sprintf("Ensuring that %s are running within %s pod", containers, addonPod))
				Expect(pod.EnsureContainersRunningInAllPods(containers, addonPod, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)).NotTo(HaveOccurred())

				// Validate CSI node windows pod
				if eng.HasWindowsAgents() && common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.18.0") {
					addonPod = fmt.Sprintf("csi-%s-node-windows", shortenedAddonName)
					containers = []string{"liveness-probe", "node-driver-registrar", shortenedAddonName}
					By(fmt.Sprintf("Ensuring that %s are running within %s pod", containers, addonPod))
					Expect(pod.EnsureContainersRunningInAllPods(containers, addonPod, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)).NotTo(HaveOccurred())
				}

				// Validate CSI snapshot controller pod
				switch addonName {
				case common.AzureDiskCSIDriverAddonName:
					if eng.AnyAgentIsLinux() && common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
						addonPod = "csi-snapshot-controller"
						containers = []string{"csi-snapshot-controller"}
						By(fmt.Sprintf("Ensuring that %s are running within %s pod", containers, addonPod))
						Expect(pod.EnsureContainersRunningInAllPods(containers, addonPod, "kube-system", kubeSystemPodsReadinessChecks, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)).NotTo(HaveOccurred())
					}
				}
			}
		})
	})

	Describe("with a windows agent pool", func() {
		It("kubelet service should be able to recover when the docker service is stopped", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				if eng.HasWindowsAgents() {
					if eng.ExpandedDefinition.Properties.WindowsProfile != nil && eng.ExpandedDefinition.Properties.WindowsProfile.GetSSHEnabled() {
						nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						simulateDockerdCrashScript := "simulate-dockerd-crash.cmd"
						err = sshConn.CopyTo(simulateDockerdCrashScript)
						Expect(err).NotTo(HaveOccurred())
						for _, n := range nodes {
							if n.IsWindows() {
								By(fmt.Sprintf("simulating docker and subsequent kubelet service crash on node: %s", n.Metadata.Name))
								err = sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+simulateDockerdCrashScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
								Expect(err).NotTo(HaveOccurred())
								simulateDockerCrashCommand := fmt.Sprintf("\"/tmp/%s\"", simulateDockerdCrashScript)
								err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, simulateDockerCrashCommand, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
								Expect(err).NotTo(HaveOccurred())
							}
						}

						log.Print("Waiting 1 minute to allow nodes to report not ready state after the crash occurred\n")
						time.Sleep(1 * time.Minute)

						for _, n := range nodes {
							if n.IsWindows() {
								By(fmt.Sprintf("restarting kubelet service on node: %s", n.Metadata.Name))
								restartKubeletCommand := fmt.Sprintf("\"Powershell Start-Service kubelet\"")
								err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, restartKubeletCommand, true, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
								Expect(err).NotTo(HaveOccurred())
							}
						}

						var expectedReadyNodes int
						if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() && !clusterAutoscalerEngaged {
							expectedReadyNodes = len(nodes)
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
				deploy, err := deployment.CreateLinuxDeployDeleteIfExists(deploymentPrefix, "library/nginx:latest", deploymentName, "default", serviceName, "", cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can create an ILB service attachment")
				sILB, err := service.CreateServiceFromFileDeleteIfExist(filepath.Join(WorkloadDir, "ingress-nginx-ilb.yaml"), serviceName+"-ilb", "default")
				Expect(err).NotTo(HaveOccurred())
				err = sILB.WaitForIngress(cfg.LBTimeout, 5*time.Second)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can create a curl pod to connect to the service")
				ilbCurlPod, err := pod.RunLinuxWithRetry("byrnedo/alpine-curl", "curl-to-ilb", "default", fmt.Sprintf("curl %s", sILB.Status.LoadBalancer.Ingress[0]["ip"]), false, 1*time.Minute, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can create an ELB service attachment")
				sELB, err := service.CreateServiceFromFileDeleteIfExist(filepath.Join(WorkloadDir, "ingress-nginx-elb.yaml"), serviceName+"-elb", "default")
				Expect(err).NotTo(HaveOccurred())
				err = sELB.WaitForIngress(cfg.LBTimeout, 5*time.Second)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can connect to the ELB service on the service IP")
				err = sELB.ValidateWithRetry("(Welcome to nginx)", 30*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can connect to the ELB service from another pod")
				elbCurlPod, err := pod.RunLinuxWithRetry("byrnedo/alpine-curl", "curl-to-elb", "default", fmt.Sprintf("curl %s", sELB.Status.LoadBalancer.Ingress[0]["ip"]), false, 1*time.Minute, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we can donwnload files through the ELB")
				pods, err := pod.GetAllByPrefixWithRetry(deploymentPrefix, "default", 3*time.Second, cfg.Timeout)
				for _, p := range pods {
					out, err := p.Exec("--", "/bin/bash", "-c", "base64 /dev/urandom | head -c 500000 | tee -a /usr/share/nginx/html/index.html > /dev/null")
					log.Printf("%s\n", string(out))
					Expect(err).NotTo(HaveOccurred())
				}
				err = sELB.ValidateWithRetry("(Welcome to nginx)", 30*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				err = sILB.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = sELB.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = ilbCurlPod.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = elbCurlPod.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = deploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("No linux agent was provisioned for this Cluster Definition")
			}
		})

		It("should be able to get nodes metrics", func() {
			err := node.TopNodesWithRetry(1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create a pv by deploying a pod that consumes a pvc", func() {
			if !util.IsUsingManagedDisks(eng.ExpandedDefinition.Properties.AgentPoolProfiles) {
				Skip("Skip PV test for clusters using unmanaged disks")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() &&
				cfg.TestPVC {
				By("Creating a persistent volume claim")
				pvcName := "azure-disk" // should be the same as in pvc-azuredisk.yaml
				pvc, err := persistentvolumeclaims.CreatePersistentVolumeClaimsFromFileWithRetry(filepath.Join(WorkloadDir, "pvc-azuredisk.yaml"), pvcName, "default", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				// Azure Disk CSI driver in zone-enabled clusters uses 'WaitForFirstConsumer' volume binding mode
				// thus, pvc won't be available until a pod consumes it
				isUsingAzureDiskCSIDriver, _ := eng.HasAddon("azuredisk-csi-driver")
				if !(isUsingAzureDiskCSIDriver && eng.ExpandedDefinition.Properties.HasZonesForAllAgentPools()) {
					ready, err := pvc.WaitOnReady("default", 5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))
				}

				By("Launching a pod using the volume claim")
				podName := "pv-pod" // should be the same as in pod-pvc.yaml
				testPod, err := pod.CreatePodFromFileWithRetry(filepath.Join(WorkloadDir, "pod-pvc.yaml"), podName, "default", 1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				ready, err := testPod.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(ready).To(Equal(true))

				By("Checking that the pod can access volume")
				valid, err := testPod.ValidatePVC("/mnt/azure", 10, 10*time.Second)
				Expect(valid).To(BeTrue())
				Expect(err).NotTo(HaveOccurred())

				// Skip label validation for Azure Disk CSI driver since it currently doesn't apply any label to PV
				if !isUsingAzureDiskCSIDriver && eng.ExpandedDefinition.Properties.HasZonesForAllAgentPools() {
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

					By("Ensuring that attached volume pv has the same zone as the zone of the node")
					nodeName := testPod.Spec.NodeName
					nodeList, err := node.GetByRegexWithRetry(nodeName, 3*time.Minute, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					nodeZone := nodeList[0].Metadata.Labels["failure-domain.beta.kubernetes.io/zone"]
					fmt.Printf("pvZone: %s\n", pvZone)
					fmt.Printf("nodeZone: %s\n", nodeZone)
					Expect(nodeZone == pvZone).To(Equal(true))
				}

				By("Cleaning up after ourselves")
				err = testPod.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				err = pvc.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})
	})

	Describe("with a GPU-enabled agent pool", func() {
		It("should be able to run a nvidia-gpu job", func() {
			if eng.ExpandedDefinition.Properties.HasNSeriesSKU() {
				j, err := job.CreateJobFromFileWithRetry(filepath.Join(WorkloadDir, "cuda-vector-add.yaml"), "cuda-vector-add", "default", 3*time.Second, cfg.Timeout)
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
				Skip("This is not a GPU-enabled cluster")
			}
		})
	})

	Describe("with a DC-series SKU agent pool", func() {
		It("should be able to run an SGX job", func() {
			if eng.ExpandedDefinition.Properties.HasDCSeriesSKU() {
				j, err := job.CreateJobFromFileWithRetry(filepath.Join(WorkloadDir, "sgx-test.yaml"), "sgx-test", "default", 3*time.Second, cfg.Timeout)
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

		It("should be able to run an SGX job with sgx-device-plugin", func() {
			if eng.ExpandedDefinition.Properties.HasDCSeriesSKU() {

				sgx_device_plugin := ""
				sgx_device_plugin_name := "sgx-device-plugin"
				sgx_device_plugin_namespace := "kube-system"
				sgx_device_plugin_label_key := "app"
				sgx_device_plugin_label_value := "sgx-device-plugin"

				if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
					sgx_device_plugin = "sgx-device-plugin.yaml"
				} else {
					sgx_device_plugin = "sgx-device-plugin-before-k8s-1-17.yaml"
				}

				_, err := daemonset.CreateDaemonsetFromFile(filepath.Join(WorkloadDir, sgx_device_plugin), sgx_device_plugin_name, sgx_device_plugin_namespace, 1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				pods, err := pod.GetAllRunningByLabelWithRetry(sgx_device_plugin_label_key, sgx_device_plugin_label_value, sgx_device_plugin_namespace, 1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(pods).NotTo(BeEmpty())

				j, err := job.CreateJobFromFileWithRetry(filepath.Join(WorkloadDir, "sgx-test-with-plugin.yaml"), "sgx-test-with-plugin", "default", 1*time.Second, cfg.Timeout)
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
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					var role string
					if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
						role = n.Metadata.Labels["kubernetes.azure.com/role"]
					} else {
						role = n.Metadata.Labels["kubernetes.io/role"]
					}
					if role == "master" {
						By("Ensuring that we get zones for each master node")
						zones := n.Metadata.Labels["failure-domain.beta.kubernetes.io/zone"]
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
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					var role string
					if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
						role = n.Metadata.Labels["kubernetes.azure.com/role"]
					} else {
						role = n.Metadata.Labels["kubernetes.io/role"]
					}
					if role == "agent" {
						By("Ensuring that we get zones for each agent node")
						zones := n.Metadata.Labels["failure-domain.beta.kubernetes.io/zone"]
						contains := strings.Contains(zones, "-")
						Expect(contains).To(Equal(true))
					}
				}
			} else {
				Skip("Availability zones was not configured for this Cluster Definition")
			}
		})
	})

	Describe("with NetworkPolicy enabled", func() {
		It("should apply various network policies and enforce access to nginx pod", func() {
			if eng.HasNetworkPolicy("calico") || eng.HasNetworkPolicy("azure") ||
				eng.HasNetworkPolicy("cilium") || eng.HasNetworkPolicy("antrea") {
				nsDev, nsProd := "development", "production"
				By("Creating development namespace")
				namespaceDev, err := namespace.CreateNamespaceDeleteIfExist(nsDev)
				Expect(err).NotTo(HaveOccurred())
				By("Creating production namespace")
				namespaceProd, err := namespace.CreateNamespaceDeleteIfExist(nsProd)
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
				frontendProdDeployment, err := deployment.CreateDeploymentFromImageWithRetry("library/nginx:latest", frontendProdDeploymentName, nsProd, "webapp", "frontend", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				frontendDevDeploymentName := fmt.Sprintf("frontend-dev-%s-%v", cfg.Name, randInt+100000)
				frontendDevDeployment, err := deployment.CreateDeploymentFromImageWithRetry("library/nginx:latest", frontendDevDeploymentName, nsDev, "webapp", "frontend", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				backendDeploymentName := fmt.Sprintf("backend-%s-%v", cfg.Name, randInt+200000)
				backendDeployment, err := deployment.CreateDeploymentFromImageWithRetry("library/nginx:latest", backendDeploymentName, nsDev, "webapp", "backend", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				nwpolicyDeploymentName := fmt.Sprintf("network-policy-%s-%v", cfg.Name, randInt+300000)
				nwpolicyDeployment, err := deployment.CreateDeploymentFromImageWithRetry("library/nginx:latest", nwpolicyDeploymentName, nsDev, "", "", 3*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				By("Ensure there is a running frontend-prod pod")
				networkpolicy.EnsureRunningPodExists(frontendProdDeploymentName, nsProd, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)

				By("Ensure there is a running frontend-dev pod")
				networkpolicy.EnsureRunningPodExists(frontendDevDeploymentName, nsDev, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)

				By("Ensure there is a running backend pod")
				networkpolicy.EnsureRunningPodExists(backendDeploymentName, nsDev, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)

				By("Ensure there is a running network-policy pod")
				networkpolicy.EnsureRunningPodExists(nwpolicyDeploymentName, nsDev, 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)

				By("Ensuring we have outbound internet access from the frontend-prod pods")
				frontendProdPods := networkpolicy.GetRunningPodsFromDeployment(frontendProdDeployment)
				networkpolicy.EnsureOutboundInternetAccess(frontendProdPods, cfg)

				By("Ensuring we have outbound internet access from the frontend-dev pods")
				frontendDevPods := networkpolicy.GetRunningPodsFromDeployment(frontendDevDeployment)
				networkpolicy.EnsureOutboundInternetAccess(frontendDevPods, cfg)

				By("Ensuring we have outbound internet access from the backend pods")
				backendPods := networkpolicy.GetRunningPodsFromDeployment(backendDeployment)
				networkpolicy.EnsureOutboundInternetAccess(backendPods, cfg)

				By("Ensuring we have outbound internet access from the network-policy pods")
				nwpolicyPods := networkpolicy.GetRunningPodsFromDeployment(nwpolicyDeployment)
				networkpolicy.EnsureOutboundInternetAccess(nwpolicyPods, cfg)

				By("Ensuring we have connectivity from network-policy pods to frontend-prod pods")
				networkpolicy.EnsureConnectivityResultBetweenPods(nwpolicyPods, frontendProdPods, validateNetworkPolicyTimeout, true)

				By("Ensuring we have connectivity from network-policy pods to backend pods")
				networkpolicy.EnsureConnectivityResultBetweenPods(nwpolicyPods, backendPods, validateNetworkPolicyTimeout, true)

				By("Applying a network policy to deny ingress access to app: webapp, role: backend pods in development namespace")
				nwpolicyName, namespace, nwpolicyFileName := "backend-deny-ingress", nsDev, "backend-policy-deny-ingress.yaml"
				networkpolicy.ApplyNetworkPolicy(nwpolicyName, namespace, nwpolicyFileName, PolicyDir)

				By("Ensuring we no longer have ingress access from the network-policy pods to backend pods")
				networkpolicy.EnsureConnectivityResultBetweenPods(nwpolicyPods, backendPods, validateNetworkPolicyTimeout, false)

				By("Cleaning up after ourselves")
				networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)

				By("Applying a network policy to deny egress access in development namespace")
				nwpolicyName, namespace, nwpolicyFileName = "backend-deny-egress", nsDev, "backend-policy-deny-egress.yaml"
				networkpolicy.ApplyNetworkPolicy(nwpolicyName, nsDev, nwpolicyFileName, PolicyDir)

				By("Ensuring we no longer have egress access from the network-policy pods to backend pods")
				networkpolicy.EnsureConnectivityResultBetweenPods(nwpolicyPods, backendPods, validateNetworkPolicyTimeout, false)
				networkpolicy.EnsureConnectivityResultBetweenPods(frontendDevPods, backendPods, validateNetworkPolicyTimeout, false)

				By("Cleaning up after ourselves")
				networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)

				if common.IsKubernetesVersionGe(eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.11.0") {

					By("Applying a network policy to allow egress access to app: webapp, role: frontend pods in any namespace from pods with app: webapp, role: backend labels in development namespace")
					nwpolicyName, namespace, nwpolicyFileName := "backend-allow-egress-pod-label", nsDev, "backend-policy-allow-egress-pod-label.yaml"
					networkpolicy.ApplyNetworkPolicy(nwpolicyName, namespace, nwpolicyFileName, PolicyDir)

					By("Ensuring we have egress access from pods with matching labels")
					networkpolicy.EnsureConnectivityResultBetweenPods(backendPods, frontendDevPods, validateNetworkPolicyTimeout, true)
					networkpolicy.EnsureConnectivityResultBetweenPods(backendPods, frontendProdPods, validateNetworkPolicyTimeout, true)

					By("Ensuring we don't have ingress access from pods without matching labels")
					networkpolicy.EnsureConnectivityResultBetweenPods(backendPods, nwpolicyPods, validateNetworkPolicyTimeout, false)

					By("Cleaning up after ourselves")
					networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)

					By("Applying a network policy to allow egress access to app: webapp, role: frontend pods from pods with app: webapp, role: backend labels in same development namespace")
					nwpolicyName, namespace, nwpolicyFileName = "backend-allow-egress-pod-namespace-label", nsDev, "backend-policy-allow-egress-pod-namespace-label.yaml"
					networkpolicy.ApplyNetworkPolicy(nwpolicyName, namespace, nwpolicyFileName, PolicyDir)

					By("Ensuring we have egress access from pods with matching labels")
					networkpolicy.EnsureConnectivityResultBetweenPods(backendPods, frontendDevPods, validateNetworkPolicyTimeout, true)

					By("Ensuring we don't have ingress access from pods without matching labels")
					networkpolicy.EnsureConnectivityResultBetweenPods(backendPods, frontendProdPods, validateNetworkPolicyTimeout, false)
					networkpolicy.EnsureConnectivityResultBetweenPods(backendPods, nwpolicyPods, validateNetworkPolicyTimeout, false)

					By("Cleaning up after ourselves")
					networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)

					By("Applying a network policy to only allow ingress access to app: webapp, role: backend pods in development namespace from pods in any namespace with the same labels")
					nwpolicyName, namespace, nwpolicyFileName = "backend-allow-ingress-pod-label", nsDev, "backend-policy-allow-ingress-pod-label.yaml"
					networkpolicy.ApplyNetworkPolicy(nwpolicyName, namespace, nwpolicyFileName, PolicyDir)

					By("Ensuring we have ingress access from pods with matching labels")
					networkpolicy.EnsureConnectivityResultBetweenPods(backendPods, backendPods, validateNetworkPolicyTimeout, true)

					By("Ensuring we don't have ingress access from pods without matching labels")
					networkpolicy.EnsureConnectivityResultBetweenPods(nwpolicyPods, backendPods, validateNetworkPolicyTimeout, false)

					By("Cleaning up after ourselves")
					networkpolicy.DeleteNetworkPolicy(nwpolicyName, namespace)

					By("Applying a network policy to only allow ingress access to app: webapp role:backends in development namespace from pods with label app:webapp, role: frontendProd within namespace with label purpose: development")
					nwpolicyName, namespace, nwpolicyFileName = "backend-policy-allow-ingress-pod-namespace-label", nsDev, "backend-policy-allow-ingress-pod-namespace-label.yaml"
					networkpolicy.ApplyNetworkPolicy(nwpolicyName, namespace, nwpolicyFileName, PolicyDir)

					By("Ensuring we don't have ingress access from role:frontend pods in production namespace")
					networkpolicy.EnsureConnectivityResultBetweenPods(frontendProdPods, backendPods, validateNetworkPolicyTimeout, false)

					By("Ensuring we have ingress access from role:frontend pods in development namespace")
					networkpolicy.EnsureConnectivityResultBetweenPods(frontendDevPods, backendPods, validateNetworkPolicyTimeout, true)

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
				Skip("Calico or Azure or Cilium or Antrea network policy was not provisioned for this Cluster Definition")
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
				err = iisService.WaitForIngress(cfg.LBTimeout, 5*time.Second)
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
					pass, err = iisPod.CheckWindowsOutboundConnection(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
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

				By("Ensuring we can donwnload files through the ELB")
				for _, iisPod := range iisPods {
					fileGenCmd := "(1..(500kb/34)).foreach({-join ('4489bfc5648d4ab58c7129a1d5f2f061') }) | Add-Content C:\\inetpub\\wwwroot\\iisstart.htm"
					out, err := iisPod.Exec("--", "powershell", "-command", fileGenCmd)
					log.Printf("%s\n", string(out))
					Expect(err).NotTo(HaveOccurred())
				}
				err = iisService.ValidateWithRetry("(IIS Windows Server)", 30*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

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
				if eng.HasNetworkPlugin(api.NetworkPluginKubenet) {
					Skip("This tests is not enabled for kubenet CNI on windows")
				}

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
				linuxNginxDeploy, err := deployment.CreateLinuxDeployDeleteIfExists(deploymentPrefix, "library/nginx:latest", nginxDeploymentName, "default", "", "", cfg.Timeout)
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
				successes, err := pod.RunCommandMultipleTimes(pod.RunWindowsPod, windowsImages.ServerCore, name, command, cfg.StabilityIterations, 1*time.Second, singleCommandTimeout, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))

				By("Connecting to Linux from Windows deployment")
				name = fmt.Sprintf("windows-2-linux-%s", cfg.Name)
				command = fmt.Sprintf("iwr -UseBasicParsing -TimeoutSec 60 %s", linuxService.Metadata.Name)
				successes, err = pod.RunCommandMultipleTimes(pod.RunWindowsPod, windowsImages.ServerCore, name, command, cfg.StabilityIterations, 1*time.Second, singleCommandTimeout, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(successes).To(Equal(cfg.StabilityIterations))

				By("Connecting to Windows from Linux deployment")
				name = fmt.Sprintf("linux-2-windows-%s", cfg.Name)
				command = fmt.Sprintf("wget %s", windowsService.Metadata.Name)
				successes, err = pod.RunCommandMultipleTimes(pod.RunLinuxPod, "alpine", name, command, cfg.StabilityIterations, 1*time.Second, singleCommandTimeout, cfg.Timeout)
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
			if eng.HasWindowsAgents() && !eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.NeedsContainerd() {
				useCloudControllerManager := to.Bool(eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager)
				if to.Bool(eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity) && useCloudControllerManager {
					Skip("cloud-controller-manager storageclass doesn't work w/ MSI")
				}
				orchestratorVersion := eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion
				if orchestratorVersion == "1.11.0" {
					// Failure in 1.11.0 - https://github.com/kubernetes/kubernetes/issues/65845, fixed in 1.11.1
					Skip("Kubernetes 1.11.0 has a known issue creating Azure PersistentVolumeClaim")
				} else if common.IsKubernetesVersionGe(orchestratorVersion, "1.8.0") {
					windowsImages, err := eng.GetWindowsTestImages()
					Expect(err).NotTo(HaveOccurred())

					iisAzurefileYaml, err := pod.ReplaceContainerImageFromFile(filepath.Join(WorkloadDir, "iis-azurefile.yaml"), windowsImages.IIS)
					Expect(err).NotTo(HaveOccurred())
					defer os.Remove(iisAzurefileYaml)

					By("Creating an AzureFile storage class")
					storageclassName := "azurefile" // should be the same as in storageclass-azurefile.yaml
					scFilename := "storageclass-azurefile.yaml"
					if useCloudControllerManager && common.IsKubernetesVersionGe(orchestratorVersion, "1.16.0") {
						scFilename = "storageclass-azurefile-external.yaml"
					}
					sc, err := storageclass.CreateStorageClassFromFileWithRetry(filepath.Join(WorkloadDir, scFilename), storageclassName, 3*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					ready, err := sc.WaitOnReady(5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))

					By("Creating a persistent volume claim")
					pvcName := "pvc-azurefile" // should be the same as in pvc-azurefile.yaml
					pvc, err := persistentvolumeclaims.CreatePVCFromFileDeleteIfExist(filepath.Join(WorkloadDir, "pvc-azurefile.yaml"), pvcName, "default", 3*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					ready, err = pvc.WaitOnReady("default", 5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					Expect(ready).To(Equal(true))

					By("Launching an IIS pod using the volume claim")
					podName := "iis-azurefile" // should be the same as in iis-azurefile.yaml
					iisPod, err := pod.CreatePodFromFileWithRetry(iisAzurefileYaml, podName, "default", 1*time.Second, cfg.Timeout)
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
		// This test is not parallelizable due to tainting nodes with NoSchedule
		It("should expect containers to be recreated after node restart", func() {
			if eng.HasWindowsAgents() {
				for _, profile := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
					if profile.IsWindows() {
						if profile.AvailabilityProfile == api.AvailabilitySet {
							Skip("AvailabilitySet is configured for this Cluster Definition")
						}
					}
				}

				if eng.HasNetworkPlugin(api.NetworkPluginKubenet) {
					Skip("This tests is not enabled for kubenet CNI on windows")
				}

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
				err = iisService.WaitForIngress(cfg.LBTimeout, 5*time.Second)
				Expect(err).NotTo(HaveOccurred())

				By("Verifying that the service is reachable and returns the default IIS start page")
				err = iisService.ValidateWithRetry("(IIS Windows Server)", sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())

				pods, err := iisDeploy.Pods()
				Expect(err).NotTo(HaveOccurred())
				nodeName := pods[0].Spec.NodeName
				ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
				defer cancel()

				By("Adding taint to all other Windows nodes")
				nodeList, err := node.Get()
				for _, n := range nodeList.Nodes {
					if n.IsWindows() && n.Metadata.Name != nodeName {
						n.AddTaint(node.Taint{Key: "key", Value: "value", Effect: "NoSchedule"})
					}
				}

				// Removing taints
				defer func(nodeList *node.List, nodeName string) {
					for _, n := range nodeList.Nodes {
						if n.IsWindows() && n.Metadata.Name != nodeName {
							n.RemoveTaint(node.Taint{Key: "key", Value: "value", Effect: "NoSchedule"})
						}
					}
				}(nodeList, nodeName)

				By("Restarting VM " + nodeName + " in resource group " + cfg.ResourceGroup)

				// Getting vmss for the vm
				vmssPage, err := azureClient.ListVirtualMachineScaleSets(ctx, cfg.ResourceGroup)
				vmssList := vmssPage.Values()

				// Name of VMSS of nodeName
				var vmssName string
				// InstanceID of VM in its VMSS
				var instanceID string
				for _, vmss := range vmssList {
					if !strings.Contains(nodeName, *vmss.Name) {
						continue
					}
					vmName := *vmss.Name + "_" + nodeName[len(nodeName)-1:]
					vmPage, err := azureClient.ListVirtualMachineScaleSetVMs(ctx, cfg.ResourceGroup, *vmss.Name)
					Expect(err).NotTo(HaveOccurred())

					vmList := vmPage.Values()
					for _, vm := range vmList {
						if vmName == *vm.Name {
							vmssName = *vmss.Name
							instanceID = *vm.InstanceID
							break
						}
					}
					if instanceID != "" {
						break
					}
				}
				// TODO refactor to remove the "compute" usage so the test can be run on Azure Stack
				instanceIDs := &compute.VirtualMachineScaleSetVMInstanceIDs{&[]string{instanceID}}
				err = azureClient.RestartVirtualMachineScaleSets(ctx, cfg.ResourceGroup, vmssName, instanceIDs)
				Expect(err).NotTo(HaveOccurred())

				//Wait for VM to come up
				time.Sleep(30 * time.Second)

				By("Verifying that the service is still reachable and returns the default IIS start page")
				err = iisService.ValidateWithRetry("(IIS Windows Server)", sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("No windows agent was provisioned for this Cluster Definition")
			}
		})

		// verifies that the pod logs continue to flow even during rotation
		// https://github.com/Azure/aks-engine/issues/3573
		It("should be able to rotate docker logs", func() {
			if !eng.HasWindowsAgents() {
				Skip("No windows agent was provisioned for this Cluster Definition")
			}

			windowsImages, err := eng.GetWindowsTestImages()
			loggingPodFile, err := pod.ReplaceContainerImageFromFile(filepath.Join(WorkloadDir, "validate-windows-logging.yaml"), windowsImages.ServerCore)
			Expect(err).NotTo(HaveOccurred())
			defer os.Remove(loggingPodFile)

			By("launching a pod that logs too much")
			podName := "validate-windows-logging" // should be the same as in iis-azurefile.yaml
			loggingPod, err := pod.CreatePodFromFileWithRetry(loggingPodFile, podName, "default", 1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			ready, err := loggingPod.WaitOnReady(sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(ready).To(Equal(true))

			By("validating the logs continue to flow")
			logsRotated, err := loggingPod.ValidateLogsRotate(20*time.Second, 2*time.Minute)
			Expect(err).NotTo(HaveOccurred())
			Expect(logsRotated).To(Equal(true))
		})

		// metrics endpoints failing in 1.18+
		// https://github.com/kubernetes/kubernetes/issues/95735
		It("windows should be able to get node metrics when high cpu", func() {
			if !eng.HasWindowsAgents() || !cfg.ValidateCPULoad {
				Skip("Will not validate effects of CPU load against nodes")
			}

			windowsImages, err := eng.GetWindowsTestImages()
			cpuConsumptionDeploymentFile, err := pod.ReplaceContainerImageFromFile(filepath.Join(WorkloadDir, "validate-windows-cpu-consumption.yaml"), windowsImages.ServerCore)
			Expect(err).NotTo(HaveOccurred())
			defer os.Remove(cpuConsumptionDeploymentFile)

			By("launching a deployment that consumes too much CPU")
			deploymentName := "validate-windows-cpu-consumption" // should be the same as in yaml
			cpuDeployment, err := deployment.CreateDeploymentFromFileWithRetry(cpuConsumptionDeploymentFile, deploymentName, "default", 1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			running, err := pod.WaitOnSuccesses(deploymentName, "default", 4, sleepBetweenRetriesWhenWaitingForPodReady, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(running).To(Equal(true))

			By("Scaling deployment to consuming allocatable")
			nodeList, err := node.GetWithRetry(1*time.Second, cfg.Timeout)
			cpuCapacity := 0
			for _, n := range nodeList {
				if n.IsWindows() {
					c, err := strconv.Atoi(n.Status.Capacity.CPU)
					Expect(err).NotTo(HaveOccurred())
					cpuCapacity = cpuCapacity + c
				}
			}

			// scale over allocatable for windows to make sure it's packed (.25 is limit on deployment)
			deployCount := int(math.Round((float64(cpuCapacity) / 0.25)))
			err = cpuDeployment.ScaleDeployment(deployCount * 2)
			Expect(err).NotTo(HaveOccurred())

			By("should be able to get nodes metrics")
			checkMetrics := func() error {
				log.Printf("running top nodes")
				err = node.TopNodes()
				return err
			}
			_, err = cpuDeployment.WaitForReplicasWithAction(deployCount, deployCount*2, 2*time.Second, cfg.Timeout, checkMetrics)
			Expect(err).NotTo(HaveOccurred())
			cpuPods, err := cpuDeployment.PodsRunning()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(cpuPods)).To(BeNumerically(">=", deployCount))

			By("should be able to get nodes metrics")
			err = node.TopNodesWithRetry(1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())

			By("Verifying pods & services can be deleted")
			err = cpuDeployment.Delete(util.DefaultDeleteRetries)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("after the cluster has been up for a while", func() {
		It("dns-liveness pod should not have any restarts", func() {
			if !cfg.RebootControlPlaneNodes {
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
			}
		})

		It("should have healthy time synchronization", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				timeSyncValidateScript := "time-sync-validate.sh"
				err = sshConn.CopyTo(timeSyncValidateScript)
				Expect(err).NotTo(HaveOccurred())
				timeSyncValidationCommand := fmt.Sprintf("\"/tmp/%s\"", timeSyncValidateScript)
				err = sshConn.Execute(timeSyncValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
						err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+timeSyncValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, timeSyncValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should be able to autoscale", func() {
			var numCoreDNSPods int
			var testCoreDNSScaleOut bool
			if eng.AnyAgentIsLinux() && eng.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs {
				// Inspired by http://blog.kubernetes.io/2016/07/autoscaling-in-kubernetes.html
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				By("Creating a php-apache deployment")
				phpApacheDeploy, err := deployment.CreateLinuxDeployIfNotExist("deis/hpa-example", longRunningApacheDeploymentName, "default", "", "", 3*time.Second, cfg.Timeout)
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
				var cpuTarget, totalMaxPods int
				if clusterAutoscalerEngaged {
					nodeList, err := node.GetWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					if hasAddon, addon := eng.HasAddon("coredns"); hasAddon {
						nodesPerReplica, _ := strconv.Atoi(addon.Config["nodes-per-replica"])
						corednsPods, err := pod.GetAllByPrefixWithRetry("coredns", "kube-system", 3*time.Second, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						corednsAutoscalerPods, err := pod.GetAllByPrefixWithRetry("coredns-autoscaler", "kube-system", 3*time.Second, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						numCoreDNSPods = len(corednsPods) - len(corednsAutoscalerPods)
						coreDNSNodesOverhead := nodesPerReplica - (len(nodeList) * numCoreDNSPods)
						var clusterAutoscalerNodesOverhead int
						for _, pool := range clusterAutoscalerAddon.Pools {
							p := eng.ExpandedDefinition.Properties.GetAgentPoolIndexByName(pool.Name)
							maxNodes, _ := strconv.Atoi(pool.Config["max-nodes"])
							if maxNodes > eng.ExpandedDefinition.Properties.AgentPoolProfiles[p].Count {
								clusterAutoscalerNodesOverhead += (maxNodes - eng.ExpandedDefinition.Properties.AgentPoolProfiles[p].Count)
							}
						}
						if coreDNSNodesOverhead >= 0 && coreDNSNodesOverhead < clusterAutoscalerNodesOverhead {
							testCoreDNSScaleOut = true
							By("Validating that coredns pods scale out with nodes")
							log.Printf("%d coredns pods before scaling out\n", numCoreDNSPods)
						}
					}
					cpuTarget = 50
					for _, profile := range eng.ExpandedDefinition.Properties.AgentPoolProfiles {
						// TODO enable cluster-autoscaler tests for Windows
						if profile.IsLinux() {
							for _, pool := range clusterAutoscalerAddon.Pools {
								if pool.Name == profile.Name {
									maxPods, _ := strconv.Atoi(profile.KubernetesConfig.KubeletConfig["--max-pods"])
									totalMaxPods += (profile.Count * maxPods)
								}
							}
						}
					}
					maxPods, _ := strconv.Atoi(eng.ExpandedDefinition.Properties.MasterProfile.KubernetesConfig.KubeletConfig["--max-pods"])
					totalMaxPods += (len(masterNodes) * maxPods)
				} else {
					cpuTarget = 50
					totalMaxPods = 10
				}
				err = phpApacheDeploy.CreateDeploymentHPADeleteIfExist(cpuTarget, 1, totalMaxPods+1)
				Expect(err).NotTo(HaveOccurred())
				h, err := hpa.Get(longRunningApacheDeploymentName, "default", 10)
				Expect(err).NotTo(HaveOccurred())

				By("Sending load to the php-apache service by creating a 3 replica deployment")
				// Launch a simple busybox pod that wget's continuously to the apache serviceto simulate load
				commandString := fmt.Sprintf("while true; do wget -q -O- http://%s.default.svc.cluster.local; done", longRunningApacheDeploymentName)
				loadTestPrefix := fmt.Sprintf("load-test-%s", cfg.Name)
				loadTestName := fmt.Sprintf("%s-%v", loadTestPrefix, r.Intn(99999))
				numLoadTestPods := 3
				if clusterAutoscalerEngaged {
					numLoadTestPods = (totalMaxPods / 2)
				}
				loadTestDeploy, err := deployment.RunLinuxDeployDeleteIfExists(loadTestPrefix, "busybox", loadTestName, "default", commandString, numLoadTestPods)
				Expect(err).NotTo(HaveOccurred())

				By("Ensuring we have more than 1 apache-php pods due to hpa enforcement")
				_, err = phpApacheDeploy.WaitForReplicas(2, -1, 5*time.Second, cfg.Timeout)
				if err != nil {
					e := h.Describe()
					Expect(e).NotTo(HaveOccurred())
				}
				Expect(err).NotTo(HaveOccurred())

				if clusterAutoscalerEngaged {
					By("Ensuring at least one more node was added by cluster-autoscaler")
					ready := node.WaitOnReadyMin(eng.NodeCount()+1, 10*time.Second, cfg.Timeout)
					Expect(ready).To(BeTrue())
					if testCoreDNSScaleOut {
						By("Ensuring at least one more coredns pod was added by coredns-autoscaler")
						d, err := deployment.GetWithRetry("coredns", "kube-system", 5*time.Second, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						numCoreDNSAutoscalerPods := 1
						_, err = d.WaitForReplicas(numCoreDNSPods+numCoreDNSAutoscalerPods+1, -1, 5*time.Second, cfg.Timeout)
						if err != nil {
							pod.PrintPodsLogs("coredns-autoscaler", "kube-system", 5*time.Second, 1*time.Minute)
						}
						Expect(err).NotTo(HaveOccurred())
						corednsPods, err := pod.GetAllByPrefixWithRetry("coredns", "kube-system", 3*time.Second, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						corednsAutoscalerPods, err := pod.GetAllByPrefixWithRetry("coredns-autoscaler", "kube-system", 3*time.Second, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						newNumCoreDNSPods := len(corednsPods) - len(corednsAutoscalerPods)
						log.Printf("%d coredns pods after scaling out\n", newNumCoreDNSPods)
						Expect(err).NotTo(HaveOccurred())
						Expect(newNumCoreDNSPods > numCoreDNSPods).To(BeTrue())
					}
				}

				By("Stopping load")
				err = loadTestDeploy.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
				var nodes []node.Node
				if clusterAutoscalerEngaged {
					By("Wait a few more mins for additional nodes to come online, so that we can more effectively calculate node count reduction")
					time.Sleep(3 * time.Minute)
					nodes, err = node.GetWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
				}

				By("Ensuring we only have 1 apache-php pod after stopping load")
				_, err = phpApacheDeploy.WaitForReplicas(-1, 1, 5*time.Second, cfg.Timeout)
				if err != nil {
					e := h.Describe()
					Expect(e).NotTo(HaveOccurred())
				}
				Expect(err).NotTo(HaveOccurred())

				if clusterAutoscalerEngaged {
					By(fmt.Sprintf("Ensuring at least one node is removed by cluster-autoscaler, waiting until we have fewer than %d nodes...", len(nodes)))
					ready := node.WaitOnReadyMax(len(nodes)-1, 30*time.Second, cfg.Timeout*2)
					Expect(ready).To(BeTrue())
				}

				By("Deleting HPA configuration")
				err = h.Delete(util.DefaultDeleteRetries)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("This flavor/version of Kubernetes doesn't support hpa autoscale")
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

		It("should have node labels specific to masters or agents", func() {
			nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
			Expect(err).NotTo(HaveOccurred())
			for _, n := range nodes {
				role := "master"
				if !strings.HasPrefix(n.Metadata.Name, fmt.Sprintf("%s-", common.LegacyControlPlaneVMPrefix)) {
					if eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
						continue
					} else {
						role = "agent"
					}
				}
				labels := n.Metadata.Labels
				Expect(labels).To(HaveKeyWithValue("kubernetes.io/role", role))
				Expect(labels).To(HaveKey(fmt.Sprintf("node-role.kubernetes.io/%s", role)))
				if role == "master" && common.IsKubernetesVersionGe(
					eng.ExpandedDefinition.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.1") {
					Expect(labels).To(HaveKeyWithValue("node.kubernetes.io/exclude-from-external-load-balancers", "true"))
					Expect(labels).To(HaveKeyWithValue("node.kubernetes.io/exclude-disruption", "true"))
				}
				// Check node labels applied by cloud-node-manager
				if hasAddon, _ := eng.HasAddon(common.CloudNodeManagerAddonName); hasAddon {
					// Can't extract zone from API model, so just ensure that zone-related labels exist
					Expect(labels).To(HaveKey("failure-domain.beta.kubernetes.io/zone"))
					Expect(labels).To(HaveKey("topology.kubernetes.io/zone"))
					region := eng.ExpandedDefinition.Location
					Expect(labels).To(HaveKeyWithValue("failure-domain.beta.kubernetes.io/region", region))
					Expect(labels).To(HaveKeyWithValue("topology.kubernetes.io/region", region))
					var instanceType string
					switch role {
					case "master":
						instanceType = eng.ExpandedDefinition.Properties.MasterProfile.VMSize
					case "agent":
						osType := api.Linux
						if n.IsWindows() {
							osType = api.Windows
						}
						instanceType = util.GetAgentVMSize(eng.ExpandedDefinition.Properties.AgentPoolProfiles, osType)
					}
					Expect(labels).To(HaveKeyWithValue("beta.kubernetes.io/instance-type", instanceType))
					Expect(labels).To(HaveKeyWithValue("node.kubernetes.io/instance-type", instanceType))
				}
			}
		})

		It("should have arc agents running", func() {
			if hasArc, _ := eng.HasAddon(common.AzureArcOnboardingAddonName); hasArc {
				By("Checking the onboarding job succeeded")
				succeeded, err := job.WaitOnSucceeded("azure-arc-onboarding", "azure-arc-onboarding", 30*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				Expect(succeeded).To(Equal(true))

				By("Checking ready status of each pod in namespace azure-arc")
				pods, err := pod.GetAll("azure-arc")
				Expect(err).NotTo(HaveOccurred())
				Expect(len(pods.Pods)).ToNot(BeZero())
				for _, currentPod := range pods.Pods {
					log.Printf("Checking %s - ready: %t, restarts: %d", currentPod.Metadata.Name, currentPod.Status.ContainerStatuses[0].Ready, currentPod.Status.ContainerStatuses[0].RestartCount)
					Expect(currentPod.Status.ContainerStatuses[0].Ready).To(BeTrue())
					tooManyRestarts := 5
					Expect(currentPod.Status.ContainerStatuses[0].RestartCount).To(BeNumerically("<", tooManyRestarts))
				}
			} else {
				Skip("Onboarding connected cluster was not requested")
			}
		})

		It("should have resilient kubelet and docker systemd services", func() {
			if cfg.BlockSSHPort {
				Skip("SSH port is blocked")
			} else if !eng.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
				Expect(err).NotTo(HaveOccurred())
				systemdValidateScript := "systemd-validate.sh"
				err = sshConn.CopyTo(systemdValidateScript)
				Expect(err).NotTo(HaveOccurred())
				systemdValidationCommand := fmt.Sprintf("/tmp/%s", systemdValidateScript)
				err = sshConn.Execute(systemdValidationCommand, false)
				Expect(err).NotTo(HaveOccurred())
				for _, n := range nodes {
					if n.IsUbuntu() && !firstMasterRegexp.MatchString(n.Metadata.Name) {
						err := sshConn.CopyToRemoteWithRetry(n.Metadata.Name, "/tmp/"+systemdValidateScript, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
						systemdValidationCommand = fmt.Sprintf("/tmp/%s", systemdValidateScript)
						err = sshConn.ExecuteRemoteWithRetry(n.Metadata.Name, systemdValidationCommand, false, sleepBetweenRetriesRemoteSSHCommand, cfg.Timeout)
						Expect(err).NotTo(HaveOccurred())
					}
				}
			} else {
				Skip("Skip per-node tests in low-priority VMSS cluster configuration scenario")
			}
		})

		It("should be able to install vmss node prototype", func() {
			if cfg.RunVMSSNodePrototype {
				if eng.ExpandedDefinition.Properties.HasVMSSAgentPool() {
					nodes, err := node.GetReadyWithRetry(1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					numAgentNodes := len(nodes) - len(masterNodes)
					By("Creating a DaemonSet with a large container")
					d, err := daemonset.CreateDaemonsetDeleteIfExists(filepath.Join(WorkloadDir, "large-container-daemonset.yaml"), "large-container-daemonset", "default", "app", "large-container-daemonset", 5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					start := time.Now()
					pods, err := pod.WaitForMinRunningByLabelWithRetry(numAgentNodes, "app", "large-container-daemonset", "default", 1*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					numLargeContainerPods := len(pods)
					Expect(pods).NotTo(BeEmpty())
					elapsed := time.Since(start)
					log.Printf("Took %s to schedule %d Pods with large containers via DaemonSet\n", elapsed, numLargeContainerPods)
					By("Choosing a target VMSS node to use as the prototype")
					var targetNode string
					for _, n := range nodes {
						if strings.Contains(n.Spec.ProviderID, "virtualMachineScaleSets") {
							targetNode = n.Metadata.Name
						}
					}
					Expect(targetNode).NotTo(BeEmpty())
					cmd := exec.Command("helm", "status", "vmss-prototype")
					out, err := cmd.CombinedOutput()
					if err == nil {
						By("Found pre-existing 'vmss-prototype' helm release, deleting it...")
						cmd := exec.Command("helm", "delete", "vmss-prototype")
						out, err := cmd.CombinedOutput()
						log.Printf("%s\n", out)
						Expect(err).NotTo(HaveOccurred())
					}
					cmd = exec.Command("helm", "upgrade", "--install",
						"--repo", "https://jackfrancis.github.io/kamino/",
						"vmss-prototype", "vmss-prototype",
						"--namespace", "default",
						"--set", "kamino.scheduleOnControlPlane=true",
						"--set", fmt.Sprintf("kamino.targetNode=%s", targetNode))
					start = time.Now()
					out, err = cmd.CombinedOutput()
					log.Printf("%s\n", out)
					Expect(err).NotTo(HaveOccurred())
					By("Ensuring that the kamino-vmss-prototype pod runs to completion")
					pods, err = pod.GetAllSucceededByLabelWithRetry("app", "kamino-vmss-prototype", "default", 5*time.Second, 120*time.Minute)
					Expect(err).NotTo(HaveOccurred())
					Expect(len(pods)).To(Equal(1))
					elapsed = time.Since(start)
					log.Printf("Took %s to run kamino-vmss-prototype Job to completion\n", elapsed)
					By("Adding one new node to ensure that daemonset with a large container gets to a Running state quickly")
					ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
					defer cancel()
					// Getting vmss for the vm
					vmssPage, err := azureClient.ListVirtualMachineScaleSets(ctx, cfg.ResourceGroup)
					vmssList := vmssPage.Values()
					// Name of VMSS of nodeName
					var vmssName string
					var vmssSku *compute.Sku
					for _, vmss := range vmssList {
						if !strings.Contains(targetNode, *vmss.Name) {
							continue
						}
						vmssName = *vmss.Name
						vmssSku = vmss.Sku
					}
					Expect(vmssName).NotTo(BeEmpty())
					start = time.Now()
					err = azureClient.SetVirtualMachineScaleSetCapacity(
						ctx,
						cfg.ResourceGroup,
						vmssName,
						compute.Sku{
							Name:     vmssSku.Name,
							Capacity: to.Int64Ptr(*vmssSku.Capacity + 1),
						},
						eng.ExpandedDefinition.Location,
					)
					By("Waiting for the new node to become Ready")
					ready := node.WaitOnReadyMin(numAgentNodes+1, 500*time.Millisecond, cfg.Timeout)
					Expect(ready).To(BeTrue())
					elapsed = time.Since(start)
					log.Printf("Took %s to add 1 node derived from peer node prototype\n", elapsed)
					By("Ensuring that we have one additional large container pod after scaling out by one")
					start = time.Now()
					_, err = pod.WaitForMinRunningByLabelWithRetry(numLargeContainerPods+1, "app", "large-container-daemonset", "default", 5*time.Second, cfg.Timeout)
					Expect(err).NotTo(HaveOccurred())
					By("Ensuring that the daemonset pod achieved a Running state in under 5 seconds")
					elapsed = time.Since(start)
					log.Printf("Took %s for large-container-daemonset pod to reach Running state on new node\n", elapsed)
					Expect(elapsed < 10*time.Second).To(BeTrue())
					By("Deleting large container DaemonSet")
					Expect(err).NotTo(HaveOccurred())
					err = d.Delete(util.DefaultDeleteRetries)
					Expect(err).NotTo(HaveOccurred())
				} else {
					Skip("no VMSS node pools")
				}
			} else {
				Skip("InstallVMSSNodePrototype disabled")
			}
		})
	})
})
