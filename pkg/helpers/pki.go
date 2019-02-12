// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"time"

	"golang.org/x/sync/errgroup"

	log "github.com/sirupsen/logrus"
)

const (
	// ValidityDuration specifies the duration an TLS certificate is valid
	ValidityDuration = time.Hour * 24 * 365 * 30
	// PkiKeySize is the size in bytes of the PKI key
	PkiKeySize = 4096
)

// PkiKeyCertPair represents an PKI public and private cert pair
type PkiKeyCertPair struct {
	CertificatePem string
	PrivateKeyPem  string
}

// CreatePkiKeyCertPair generates a pair of PKI certificate and private key
func CreatePkiKeyCertPair(commonName string) (*PkiKeyCertPair, error) {
	caCertificate, caPrivateKey, err := createCertificate(commonName, nil, nil, false, false, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	caPair := &PkiKeyCertPair{CertificatePem: string(certificateToPem(caCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(caPrivateKey))}
	return caPair, nil
}

// CreatePki creates PKI certificates
func CreatePki(extraFQDNs []string, extraIPs []net.IP, clusterDomain string, caPair *PkiKeyCertPair, masterCount int) (*PkiKeyCertPair, *PkiKeyCertPair, *PkiKeyCertPair, *PkiKeyCertPair, *PkiKeyCertPair, []*PkiKeyCertPair, error) {
	start := time.Now()
	defer func(s time.Time) {
		log.Debugf("pki: PKI asset creation took %s", time.Since(s))
	}(start)
	extraFQDNs = append(extraFQDNs, fmt.Sprintf("kubernetes"))
	extraFQDNs = append(extraFQDNs, fmt.Sprintf("kubernetes.default"))
	extraFQDNs = append(extraFQDNs, fmt.Sprintf("kubernetes.default.svc"))
	extraFQDNs = append(extraFQDNs, fmt.Sprintf("kubernetes.default.svc.%s", clusterDomain))
	extraFQDNs = append(extraFQDNs, fmt.Sprintf("kubernetes.kube-system"))
	extraFQDNs = append(extraFQDNs, fmt.Sprintf("kubernetes.kube-system.svc"))
	extraFQDNs = append(extraFQDNs, fmt.Sprintf("kubernetes.kube-system.svc.%s", clusterDomain))

	var (
		caCertificate         *x509.Certificate
		caPrivateKey          *rsa.PrivateKey
		apiServerCertificate  *x509.Certificate
		apiServerPrivateKey   *rsa.PrivateKey
		clientCertificate     *x509.Certificate
		clientPrivateKey      *rsa.PrivateKey
		kubeConfigCertificate *x509.Certificate
		kubeConfigPrivateKey  *rsa.PrivateKey
		etcdServerCertificate *x509.Certificate
		etcdServerPrivateKey  *rsa.PrivateKey
		etcdClientCertificate *x509.Certificate
		etcdClientPrivateKey  *rsa.PrivateKey
		etcdPeerCertPairs     []*PkiKeyCertPair
	)
	var group errgroup.Group

	var err error
	caCertificate, err = pemToCertificate(caPair.CertificatePem)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	caPrivateKey, err = pemToKey(caPair.PrivateKeyPem)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	group.Go(func() (err error) {
		apiServerCertificate, apiServerPrivateKey, err = createCertificate("apiserver", caCertificate, caPrivateKey, false, true, extraFQDNs, extraIPs, nil)
		return err
	})

	group.Go(func() (err error) {
		organization := make([]string, 1)
		organization[0] = "system:masters"
		clientCertificate, clientPrivateKey, err = createCertificate("client", caCertificate, caPrivateKey, false, false, nil, nil, organization)
		return err
	})

	group.Go(func() (err error) {
		organization := make([]string, 1)
		organization[0] = "system:masters"
		kubeConfigCertificate, kubeConfigPrivateKey, err = createCertificate("client", caCertificate, caPrivateKey, false, false, nil, nil, organization)
		return err
	})

	group.Go(func() (err error) {
		etcdServerCertificate, etcdServerPrivateKey, err = createCertificate("etcdserver", caCertificate, caPrivateKey, true, true, nil, extraIPs, nil)
		return err
	})

	group.Go(func() (err error) {
		etcdClientCertificate, etcdClientPrivateKey, err = createCertificate("etcdclient", caCertificate, caPrivateKey, true, false, nil, extraIPs, nil)
		return err
	})

	etcdPeerCertPairs = make([]*PkiKeyCertPair, masterCount)
	for i := 0; i < masterCount; i++ {
		i := i
		group.Go(func() (err error) {
			etcdPeerCertificate, etcdPeerPrivateKey, err := createCertificate("etcdpeer", caCertificate, caPrivateKey, true, false, nil, extraIPs, nil)
			etcdPeerCertPairs[i] = &PkiKeyCertPair{CertificatePem: string(certificateToPem(etcdPeerCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(etcdPeerPrivateKey))}
			return err
		})
	}

	if err := group.Wait(); err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return &PkiKeyCertPair{CertificatePem: string(certificateToPem(apiServerCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(apiServerPrivateKey))},
		&PkiKeyCertPair{CertificatePem: string(certificateToPem(clientCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(clientPrivateKey))},
		&PkiKeyCertPair{CertificatePem: string(certificateToPem(kubeConfigCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(kubeConfigPrivateKey))},
		&PkiKeyCertPair{CertificatePem: string(certificateToPem(etcdServerCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(etcdServerPrivateKey))},
		&PkiKeyCertPair{CertificatePem: string(certificateToPem(etcdClientCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(etcdClientPrivateKey))},
		etcdPeerCertPairs,
		nil
}

func createCertificate(commonName string, caCertificate *x509.Certificate, caPrivateKey *rsa.PrivateKey, isEtcd bool, isServer bool, extraFQDNs []string, extraIPs []net.IP, organization []string) (*x509.Certificate, *rsa.PrivateKey, error) {
	var err error

	isCA := (caCertificate == nil)

	now := time.Now()

	template := x509.Certificate{
		Subject:   pkix.Name{CommonName: commonName},
		NotBefore: now,
		NotAfter:  now.Add(ValidityDuration),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	if organization != nil {
		template.Subject.Organization = organization
	}

	if isCA {
		template.KeyUsage |= x509.KeyUsageCertSign
		template.IsCA = isCA
	} else if isEtcd {
		if commonName == "etcdServer" {
			template.IPAddresses = extraIPs
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
		} else if commonName == "etcdClient" {
			template.IPAddresses = extraIPs
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
		} else {
			template.IPAddresses = extraIPs
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
		}
	} else if isServer {
		template.DNSNames = extraFQDNs
		template.IPAddresses = extraIPs
		template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
	} else {
		template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
	}

	snMax := new(big.Int).Lsh(big.NewInt(1), 128)
	template.SerialNumber, err = rand.Int(rand.Reader, snMax)
	if err != nil {
		return nil, nil, err
	}

	privateKey, _ := rsa.GenerateKey(rand.Reader, PkiKeySize)

	var privateKeyToUse *rsa.PrivateKey
	var certificateToUse *x509.Certificate
	if !isCA {
		privateKeyToUse = caPrivateKey
		certificateToUse = caCertificate
	} else {
		privateKeyToUse = privateKey
		certificateToUse = &template
	}

	certDerBytes, err := x509.CreateCertificate(rand.Reader, &template, certificateToUse, &privateKey.PublicKey, privateKeyToUse)
	if err != nil {
		return nil, nil, err
	}

	certificate, err := x509.ParseCertificate(certDerBytes)
	if err != nil {
		return nil, nil, err
	}

	return certificate, privateKey, nil
}

func certificateToPem(derBytes []byte) []byte {
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	pemBuffer := bytes.Buffer{}
	pem.Encode(&pemBuffer, pemBlock)

	return pemBuffer.Bytes()
}

func privateKeyToPem(privateKey *rsa.PrivateKey) []byte {
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	pemBuffer := bytes.Buffer{}
	pem.Encode(&pemBuffer, pemBlock)

	return pemBuffer.Bytes()
}

func pemToCertificate(raw string) (*x509.Certificate, error) {
	cpb, _ := pem.Decode([]byte(raw))
	if cpb == nil {
		return nil, errors.New("The raw pem is not a valid PEM formatted block")
	}
	return x509.ParseCertificate(cpb.Bytes)
}

func pemToKey(raw string) (*rsa.PrivateKey, error) {
	kpb, _ := pem.Decode([]byte(raw))
	if kpb == nil {
		return nil, errors.New("The raw pem is not a valid PEM formatted block")
	}
	return x509.ParsePKCS1PrivateKey(kpb.Bytes)
}
