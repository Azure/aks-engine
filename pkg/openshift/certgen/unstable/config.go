// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package unstable

import (
	"crypto/rsa"
	"crypto/x509"
	"math/big"
	"net"
	"sync"

	"github.com/Azure/aks-engine/pkg/openshift/filesystem"
)

// Config represents an OpenShift configuration
type Config struct {
	ExternalMasterHostname  string
	serial                  serial
	cas                     map[string]CertAndKey
	AuthSecret              string
	EncSecret               string
	Master                  *Master
	Bootstrap               KubeConfig
	ClusterUsername         string
	ClusterPassword         string
	EnableAADAuthentication bool
	AzureConfig             AzureConfig
}

// AzureConfig represents the azure.conf configuration
type AzureConfig struct {
	TenantID                   string
	SubscriptionID             string
	AADClientID                string
	AADClientSecret            string
	ResourceGroup              string
	Location                   string
	SecurityGroupName          string
	PrimaryAvailabilitySetName string
}

// Master represents an OpenShift master configuration
type Master struct {
	Hostname string
	IPs      []net.IP
	Port     int16

	certs       map[string]CertAndKey
	etcdcerts   map[string]CertAndKey
	kubeconfigs map[string]KubeConfig
}

// CertAndKey is a certificate and key
type CertAndKey struct {
	cert *x509.Certificate
	key  *rsa.PrivateKey
}

type serial struct {
	m sync.Mutex
	i int64
}

func (s *serial) Get() *big.Int {
	s.m.Lock()
	defer s.m.Unlock()

	s.i++
	return big.NewInt(s.i)
}

// WriteMaster writes the config files for a Master node to a Filesystem.
func (c *Config) WriteMaster(fs filesystem.Writer) error {
	err := c.WriteMasterCerts(fs)
	if err != nil {
		return err
	}

	err = c.WriteMasterKeypair(fs)
	if err != nil {
		return err
	}

	err = c.WriteMasterKubeConfigs(fs)
	if err != nil {
		return err
	}

	err = c.WriteMasterFiles(fs)
	if err != nil {
		return err
	}

	err = c.WriteBootstrapCerts(fs)
	if err != nil {
		return err
	}

	return c.WriteNodeFiles(fs)
}

// WriteNode writes the config files for bootstrapping a node to a Filesystem.
func (c *Config) WriteNode(fs filesystem.Writer) error {
	err := c.WriteBootstrapCerts(fs)
	if err != nil {
		return err
	}

	err = c.WriteBootstrapKubeConfig(fs)
	if err != nil {
		return err
	}

	return c.WriteNodeFiles(fs)
}
