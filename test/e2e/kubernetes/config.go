package kubernetes

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
)

// Config represents a kubernetes config object
type Config struct {
	Clusters []Cluster `json:"clusters"`
}

// Cluster contains the name and the cluster info
type Cluster struct {
	Name        string      `json:"name"`
	ClusterInfo ClusterInfo `json:"cluster"`
}

// ClusterInfo holds the server and cert
type ClusterInfo struct {
	Server string `json:"server"`
}

// GetConfig returns a Config value representing the current kubeconfig
func GetConfig() (*Config, error) {
	cmd := exec.Command("kubectl", "config", "view", "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl config view':%s\n", err)
		return nil, err
	}
	c := Config{}
	err = json.Unmarshal(out, &c)
	if err != nil {
		log.Printf("Error unmarshalling config json:%s\n", err)
	}
	return &c, nil
}

// GetServerName returns the server for the given config in an sshable format
func (c *Config) GetServerName() string {
	s := c.Clusters[0].ClusterInfo.Server
	return strings.Split(s, "://")[1]
}
