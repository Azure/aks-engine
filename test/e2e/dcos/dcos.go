// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package dcos

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/engine"
	"github.com/Azure/aks-engine/test/e2e/remote"
	"github.com/pkg/errors"
)

// Cluster holds information on how to communicate with the the dcos instances
type Cluster struct {
	AdminUsername string
	AgentFQDN     string
	Connection    *remote.Connection
}

// Node represents a node object returned from querying the v1/nodes api
type Node struct {
	Host   string `json:"host_ip"`
	Health int    `json:"health"`
	Role   string `json:"role"`
}

// List holds a slice of nodes
type List struct {
	Nodes []Node `json:"nodes"`
}

// Version holds response from calling http://localhost:80/dcos-metadata/dcos-version.json
type Version struct {
	Version string `json:"version"`
}

// MarathonApp is the parent struct for a marathon app declared as json
type MarathonApp struct {
	ID                    string              `json:"id"`
	Instances             int                 `json:"instances"`
	CPUS                  float64             `json:"cpus"`
	Memory                int                 `json:"mem"`
	Disk                  int                 `json:"disk"`
	GPUS                  int                 `json:"gpus"`
	BackoffSeconds        int                 `json:"backoffSeconds"`
	BackoffFactor         float64             `json:"backoffFactor"`
	MaxLaunchDelaySeconds int                 `json:"maxLaunchDelaySeconds"`
	RequirePorts          bool                `json:"requirePorts"`
	KillSelection         string              `json:"killSelection"`
	TaskHealthy           int                 `json:"tasksHealthy"`
	TaskRunning           int                 `json:"tasksRunning"`
	TaskStaged            int                 `json:"tasksStaged"`
	TaskUnhealthy         int                 `json:"tasksUnhealthy"`
	Container             Container           `json:"container"`
	HealthChecks          []HealthCheck       `json:"healthChecks"`
	UpgradeStrategy       UpgradeStrategy     `json:"upgradeStrategy"`
	UnreachableStrategy   UnreachableStrategy `json:"unreachableStrategy"`
	AcceptedResourceRoles []string            `json:"acceptedResourceRoles"`
	Labels                map[string]string   `json:"labels"`
}

// Container holds information about the type of container being deployed
type Container struct {
	Type   string `json:"type"`
	Docker Docker `json:"docker"`
}

// Docker tells what image is being deployed and its port mappings
type Docker struct {
	Image          string    `json:"image"`
	Network        string    `json:"network"`
	Priviledged    bool      `json:"priviledged"`
	ForcePullImage bool      `json:"forcePullImage"`
	PortMappings   []PortMap `json:"portMappings"`
}

// PortMap is how the ports are exposed to the system and container
type PortMap struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	HostPort      int    `json:"hostPort"`
	ServicePort   int    `json:"servicePort"`
	Protocol      string `json:"protocol"`
}

// HealthCheck contains the information needed to tell DCOS how to health check a given app
type HealthCheck struct {
	GracePeriodSeconds     int    `json:"gracePeriodSeconds"`
	IntervalSeconds        int    `json:"intervalSeconds"`
	TimeoutSeconds         int    `json:"timeoutSeconds"`
	MaxConsecutiveFailures int    `json:"maxConsecutiveFailures"`
	PortIndex              int    `json:"portIndex"`
	Path                   string `json:"path"`
	Protocol               string `json:"protocol"`
	IgnoreHTTP1xx          bool   `json:"ignoreHttp1xx"`
}

// UpgradeStrategy holds how many instances can be up or down during an upgrade
type UpgradeStrategy struct {
	MinimumHealthCapacity int `json:"minimumHealthCapacity"`
	MaximumOverCapacity   int `json:"maximumOverCapacity"`
}

// UnreachableStrategy tells how long to wait if an instance isnt reachable
type UnreachableStrategy struct {
	InactiveAfterSeconds int `json:"inactiveAfterSeconds"`
	ExpungeAfterSeconds  int `json:"expungeAfterSeconds"`
}

// NewCluster returns a new cluster struct
func NewCluster(cfg *config.Config, eng *engine.Engine) (*Cluster, error) {
	conn, err := remote.NewConnection(fmt.Sprintf("%s.%s.cloudapp.azure.com", cfg.Name, cfg.Location), "2200", eng.ClusterDefinition.Properties.LinuxProfile.AdminUsername, cfg.GetSSHKeyPath())
	if err != nil {
		return nil, err
	}
	return &Cluster{
		AdminUsername: eng.ClusterDefinition.Properties.LinuxProfile.AdminUsername,
		AgentFQDN:     fmt.Sprintf("%s-0.%s.cloudapp.azure.com", cfg.Name, cfg.Location),
		Connection:    conn,
	}, nil
}

// InstallDCOSClient will download and place in the path the dcos client
func (c *Cluster) InstallDCOSClient() error {

	out, err := c.Connection.Execute("curl -O https://dcos-mirror.azureedge.net/binaries/cli/linux/x86-64/dcos-1.10/dcos")
	if err != nil {
		log.Printf("Error downloading DCOS cli:%s\n", err)
		log.Printf("Output:%s\n", out)
		return err
	}
	out, err = c.Connection.Execute("chmod a+x dcos")
	if err != nil {
		log.Printf("Error trying to chmod +x the dcos cli:%s\n", err)
		log.Printf("Output:%s\n", out)
		return err
	}
	out, err = c.Connection.Execute("./dcos cluster setup http://localhost:80")
	if err != nil {
		log.Printf("Error while trying dcos cluster setup:%s\n", err)
		log.Printf("Output:%s\n", out)
		return err
	}
	return nil
}

// WaitForNodes will return an false if the nodes never become healthy
func (c *Cluster) WaitForNodes(nodeCount int, sleep, duration time.Duration) bool {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for nodes to become ready", duration.String())
			default:
				nodes, err := c.GetNodes()
				ready := true
				if err == nil {
					for _, n := range nodes {
						if n.Health != 0 {
							ready = false
						}
					}
				}
				if ready {
					readyCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case <-errCh:
			return false
		case ready := <-readyCh:
			return ready
		}
	}
}

// GetNodes will return a []Node for a given cluster
func (c *Cluster) GetNodes() ([]Node, error) {
	out, err := c.Connection.Execute("curl -s http://localhost:1050/system/health/v1/nodes")
	if err != nil {
		return nil, err
	}
	list := List{}
	err = json.Unmarshal(out, &list)
	if err != nil {
		log.Printf("Error while trying to unmarshall json:%s\n JSON:%s\n", err, out)
		return nil, err
	}
	return list.Nodes, nil
}

// NodeCount will return the node count for a dcos cluster
func (c *Cluster) NodeCount() (int, error) {
	nodes, err := c.GetNodes()
	if err != nil {
		return 0, err
	}
	return len(nodes), nil
}

// AppCount will determine the number of apps installed
func (c *Cluster) AppCount() (int, error) {
	count := 0
	out, err := c.Connection.Execute("./dcos marathon app list | sed -n '1!p' | wc -l")
	if err != nil {
		log.Printf("Error trying to fetch app count from dcos:%s\n", out)
		return count, err
	}

	count, err = strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		log.Printf("Error trying to parse output to int:%s\n", err)
		return count, err
	}
	// We should not count the marathon-lb as part of the installed app count
	if count > 0 {
		count = count - 1
	}
	return count, nil
}

// Version will return the node count for a dcos cluster
func (c *Cluster) Version() (string, error) {
	out, err := c.Connection.Execute("curl -s http://localhost:80/dcos-metadata/dcos-version.json")
	if err != nil {
		log.Printf("Error while executing connection:%s\n", err)
	}
	version := Version{}
	err = json.Unmarshal(out, &version)
	if err != nil {
		log.Printf("Error while trying to unmarshall json:%s\n JSON:%s\n", err, out)
		return "", err
	}
	return version.Version, nil
}

// InstallMarathonApp will send the marathon.json file to the remote server and install it using the dcos cli
func (c *Cluster) InstallMarathonApp(filepath string, sleep, duration time.Duration) (int, error) {
	port := 0
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("Error while trying to read marathon definition at (%s):%s\n", filepath, err)
		return 0, err
	}

	appCount, err := c.AppCount()
	if err != nil {
		return port, err
	}
	var app MarathonApp
	json.Unmarshal(contents, &app)
	app.ID = fmt.Sprintf("%s-%v", app.ID, appCount)
	for idx, pm := range app.Container.Docker.PortMappings {
		if pm.Name == "default" {
			port = pm.ServicePort + appCount
			app.Container.Docker.PortMappings[idx].ServicePort = port
		}
	}

	appJSON, err := json.Marshal(app)
	if err != nil {
		log.Printf("Error marshalling json:%s\n", err)
		return port, err
	}

	fileName := fmt.Sprintf("marathon.%v.json", appCount)
	err = c.Connection.Write(strconv.Quote(string(appJSON)), fileName)
	if err != nil {
		return port, err
	}

	if !c.AppExists(app.ID) {
		_, err = c.Connection.Execute(fmt.Sprintf("./dcos marathon app add %s", fileName))
		if err != nil {
			return 0, err
		}
		ready := c.WaitOnReady(app.ID, sleep, duration)
		if !ready {
			return 0, errors.Errorf("App %s was never installed", app.ID)
		}
	}
	return port, nil
}

// InstallMarathonLB will setup a loadbalancer if one has not been created
func (c *Cluster) InstallMarathonLB() error {
	if !c.PackageExists("marathon-lb") {
		_, err := c.Connection.Execute("./dcos package install marathon-lb --yes")
		if err != nil {
			return err
		}
	}
	return nil
}

// AppExists queries the marathon app list to see if an app exists for a given path
func (c *Cluster) AppExists(path string) bool {
	cmd := fmt.Sprintf("./dcos marathon app list | grep %s", path)
	_, err := c.Connection.Execute(cmd)
	return err == nil
}

// AppHealthy returns true if the app is deployed and healthy
func (c *Cluster) AppHealthy(path string) bool {
	cmd := fmt.Sprintf("./dcos marathon app show %s", path)
	out, err := c.Connection.Execute(cmd)
	if err != nil {
		return false
	}

	var app MarathonApp
	json.Unmarshal(out, &app)
	return app.Instances == app.TaskHealthy
}

// PackageExists retruns true if the package name is found when doing dcos package list
func (c *Cluster) PackageExists(name string) bool {
	cmd := fmt.Sprintf("./dcos package list | grep %s", name)
	_, err := c.Connection.Execute(cmd)
	return err == nil
}

// WaitOnReady will block until app is in ready state
func (c *Cluster) WaitOnReady(path string, sleep, duration time.Duration) bool {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for app (%s) to become ready", duration.String(), path)
			default:
				if c.AppExists(path) && c.AppHealthy(path) {
					time.Sleep(sleep)
					readyCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case <-errCh:
			return false
		case ready := <-readyCh:
			return ready
		}
	}
}
