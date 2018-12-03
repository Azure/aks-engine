// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package release39

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/aks-engine/pkg/openshift/filesystem"
	yaml "gopkg.in/yaml.v2"
)

// KubeConfig represents a kubeconfig
type KubeConfig struct {
	APIVersion     string                 `yaml:"apiVersion,omitempty"`
	Kind           string                 `yaml:"kind,omitempty"`
	Clusters       []Cluster              `yaml:"clusters,omitempty"`
	Contexts       []Context              `yaml:"contexts,omitempty"`
	CurrentContext string                 `yaml:"current-context,omitempty"`
	Preferences    map[string]interface{} `yaml:"preferences,omitempty"`
	Users          []User                 `yaml:"users,omitempty"`
}

// Cluster represents a kubeconfig cluster
type Cluster struct {
	Name    string      `yaml:"name,omitempty"`
	Cluster ClusterInfo `yaml:"cluster,omitempty"`
}

// ClusterInfo represents a kubeconfig clusterinfo
type ClusterInfo struct {
	Server                   string `yaml:"server,omitempty"`
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
}

// Context represents a kubeconfig context
type Context struct {
	Name    string      `yaml:"name,omitempty"`
	Context ContextInfo `yaml:"context,omitempty"`
}

// ContextInfo represents a kubeconfig contextinfo
type ContextInfo struct {
	Cluster   string `yaml:"cluster,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
	User      string `yaml:"user,omitempty"`
}

// User represents a kubeconfig user
type User struct {
	Name string   `yaml:"name,omitempty"`
	User UserInfo `yaml:"user,omitempty"`
}

// UserInfo represents a kubeconfig userinfo
type UserInfo struct {
	ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
	ClientKeyData         string `yaml:"client-key-data,omitempty"`
}

// PrepareMasterKubeConfigs creates the master kubeconfigs
func (c *Config) PrepareMasterKubeConfigs() error {
	endpoint := fmt.Sprintf("%s:%d", c.Master.Hostname, c.Master.Port)
	endpointName := strings.Replace(endpoint, ".", "-", -1)

	externalEndpoint := fmt.Sprintf("%s:%d", c.ExternalMasterHostname, c.Master.Port)
	externalEndpointName := strings.Replace(externalEndpoint, ".", "-", -1)

	localhostEndpoint := fmt.Sprintf("localhost:%d", c.Master.Port)
	localhostEndpointName := strings.Replace(localhostEndpoint, ".", "-", -1)

	cacert, err := certAsBytes(c.cas["etc/origin/master/ca"].cert)
	if err != nil {
		return err
	}
	admincert, err := certAsBytes(c.Master.certs["etc/origin/master/admin"].cert)
	if err != nil {
		return err
	}
	adminkey, err := privateKeyAsBytes(c.Master.certs["etc/origin/master/admin"].key)
	if err != nil {
		return err
	}
	mastercert, err := certAsBytes(c.Master.certs["etc/origin/master/openshift-master"].cert)
	if err != nil {
		return err
	}
	masterkey, err := privateKeyAsBytes(c.Master.certs["etc/origin/master/openshift-master"].key)
	if err != nil {
		return err
	}
	aggregatorcert, err := certAsBytes(c.Master.certs["etc/origin/master/aggregator-front-proxy"].cert)
	if err != nil {
		return err
	}
	aggregatorkey, err := privateKeyAsBytes(c.Master.certs["etc/origin/master/aggregator-front-proxy"].key)
	if err != nil {
		return err
	}

	c.Master.kubeconfigs = map[string]KubeConfig{
		"etc/origin/master/admin.kubeconfig": {
			APIVersion: "v1",
			Kind:       "Config",
			Clusters: []Cluster{
				{
					Name: externalEndpointName,
					Cluster: ClusterInfo{
						Server:                   fmt.Sprintf("https://%s", externalEndpoint),
						CertificateAuthorityData: base64.StdEncoding.EncodeToString(cacert),
					},
				},
			},
			Contexts: []Context{
				{
					Name: fmt.Sprintf("default/%s/system:admin", externalEndpointName),
					Context: ContextInfo{
						Cluster:   externalEndpointName,
						Namespace: "default",
						User:      fmt.Sprintf("system:admin/%s", externalEndpointName),
					},
				},
			},
			CurrentContext: fmt.Sprintf("default/%s/system:admin", externalEndpointName),
			Users: []User{
				{
					Name: fmt.Sprintf("system:admin/%s", externalEndpointName),
					User: UserInfo{
						ClientCertificateData: base64.StdEncoding.EncodeToString(admincert),
						ClientKeyData:         base64.StdEncoding.EncodeToString(adminkey),
					},
				},
			},
		},
		"etc/origin/master/aggregator-front-proxy.kubeconfig": {
			APIVersion: "v1",
			Kind:       "Config",
			Clusters: []Cluster{
				{
					Name: localhostEndpointName,
					Cluster: ClusterInfo{
						Server:                   fmt.Sprintf("https://%s", localhostEndpoint),
						CertificateAuthorityData: base64.StdEncoding.EncodeToString(cacert),
					},
				},
			},
			Contexts: []Context{
				{
					Name: fmt.Sprintf("default/%s/aggregator-front-proxy", localhostEndpointName),
					Context: ContextInfo{
						Cluster:   localhostEndpointName,
						Namespace: "default",
						User:      fmt.Sprintf("aggregator-front-proxy/%s", localhostEndpointName),
					},
				},
			},
			CurrentContext: fmt.Sprintf("default/%s/aggregator-front-proxy", localhostEndpointName),
			Users: []User{
				{
					Name: fmt.Sprintf("aggregator-front-proxy/%s", localhostEndpointName),
					User: UserInfo{
						ClientCertificateData: base64.StdEncoding.EncodeToString(aggregatorcert),
						ClientKeyData:         base64.StdEncoding.EncodeToString(aggregatorkey),
					},
				},
			},
		},
		"etc/origin/master/openshift-master.kubeconfig": {
			APIVersion: "v1",
			Kind:       "Config",
			Clusters: []Cluster{
				{
					Name: endpointName,
					Cluster: ClusterInfo{
						Server:                   fmt.Sprintf("https://%s", endpoint),
						CertificateAuthorityData: base64.StdEncoding.EncodeToString(cacert),
					},
				},
			},
			Contexts: []Context{
				{
					Name: fmt.Sprintf("default/%s/system:openshift-master", endpointName),
					Context: ContextInfo{
						Cluster:   endpointName,
						Namespace: "default",
						User:      fmt.Sprintf("system:openshift-master/%s", endpointName),
					},
				},
			},
			CurrentContext: fmt.Sprintf("default/%s/system:openshift-master", endpointName),
			Users: []User{
				{
					Name: fmt.Sprintf("system:openshift-master/%s", endpointName),
					User: UserInfo{
						ClientCertificateData: base64.StdEncoding.EncodeToString(mastercert),
						ClientKeyData:         base64.StdEncoding.EncodeToString(masterkey),
					},
				},
			},
		},
	}

	c.Master.kubeconfigs["etc/origin/node/bootstrap.kubeconfig"] = c.Master.kubeconfigs["etc/origin/master/admin.kubeconfig"]

	return nil
}

// PrepareBootstrapKubeConfig creates the node bootstrap kubeconfig
func (c *Config) PrepareBootstrapKubeConfig() error {
	ep := fmt.Sprintf("%s:%d", c.ExternalMasterHostname, c.Master.Port)
	epName := strings.Replace(ep, ".", "-", -1)

	cacert, err := certAsBytes(c.cas["etc/origin/master/ca"].cert)
	if err != nil {
		return err
	}

	bootstrapCert, err := certAsBytes(c.Master.certs["etc/origin/master/node-bootstrapper"].cert)
	if err != nil {
		return err
	}
	bootstrapKey, err := privateKeyAsBytes(c.Master.certs["etc/origin/master/node-bootstrapper"].key)
	if err != nil {
		return err
	}

	c.Bootstrap = KubeConfig{
		APIVersion: "v1",
		Kind:       "Config",
		Clusters: []Cluster{
			{
				Name: epName,
				Cluster: ClusterInfo{
					Server:                   fmt.Sprintf("https://%s", ep),
					CertificateAuthorityData: base64.StdEncoding.EncodeToString(cacert),
				},
			},
		},
		Contexts: []Context{
			{
				Name: fmt.Sprintf("default/%s/system:serviceaccount:openshift-infra:node-bootstrapper", epName),
				Context: ContextInfo{
					Cluster:   epName,
					Namespace: "default",
					User:      fmt.Sprintf("system:serviceaccount:openshift-infra:node-bootstrapper/%s", epName),
				},
			},
		},
		CurrentContext: fmt.Sprintf("default/%s/system:serviceaccount:openshift-infra:node-bootstrapper", epName),
		Users: []User{
			{
				Name: fmt.Sprintf("system:serviceaccount:openshift-infra:node-bootstrapper/%s", epName),
				User: UserInfo{
					ClientCertificateData: base64.StdEncoding.EncodeToString(bootstrapCert),
					ClientKeyData:         base64.StdEncoding.EncodeToString(bootstrapKey),
				},
			},
		},
	}

	return nil
}

// WriteMasterKubeConfigs writes the master kubeconfigs
func (c *Config) WriteMasterKubeConfigs(fs filesystem.Writer) error {
	for filename, kubeconfig := range c.Master.kubeconfigs {
		b, err := yaml.Marshal(&kubeconfig)
		if err != nil {
			return err
		}
		err = fs.WriteFile(filename, b, filesystem.Fileinfo{User: "root", Group: "root", Mode: os.FileMode(0600)})
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteBootstrapKubeConfig writes the node bootstrap kubeconfig
func (c *Config) WriteBootstrapKubeConfig(fs filesystem.Writer) error {
	b, err := yaml.Marshal(&c.Bootstrap)
	if err != nil {
		return err
	}

	return fs.WriteFile("etc/origin/node/bootstrap.kubeconfig", b, filesystem.Fileinfo{User: "root", Group: "root", Mode: os.FileMode(0600)})
}
