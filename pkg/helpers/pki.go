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
)

// PkiParams is used when we create Pki
type PkiParams struct {
	ExtraFQDNs    []string
	ExtraIPs      []net.IP
	ClusterDomain string
	CaPair        *PkiKeyCertPair
	MasterCount   int
	PkiKeySize    int
}

// PkiKeyCertPairParams is the params when we create the pki key cert pair.
type PkiKeyCertPairParams struct {
	CommonName string
	PkiKeySize int
}

// PkiKeyCertPair represents an PKI public and private cert pair
type PkiKeyCertPair struct {
	CertificatePem string
	PrivateKeyPem  string
}

// CreatePkiKeyCertPair generates a pair of PKI certificate and private key
func CreatePkiKeyCertPair(params PkiKeyCertPairParams) (*PkiKeyCertPair, error) {
	certPram := certParams{
		commonName:    params.CommonName,
		caCertificate: nil,
		caPrivateKey:  nil,
		isEtcd:        false,
		isServer:      false,
		extraFQDNs:    nil,
		extraIPs:      nil,
		organization:  nil,
		keySize:       params.PkiKeySize,
	}
	caCertificate, caPrivateKey, err := createCertificate(certPram)
	if err != nil {
		return nil, err
	}
	caPair := &PkiKeyCertPair{CertificatePem: string(certificateToPem(caCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(caPrivateKey))}
	return caPair, nil
}

// CreatePki creates PKI certificates
func CreatePki(pkiParams PkiParams) (*PkiKeyCertPair, *PkiKeyCertPair, *PkiKeyCertPair, *PkiKeyCertPair, *PkiKeyCertPair, []*PkiKeyCertPair, error) {
	start := time.Now()
	defer func(s time.Time) {
		log.Debugf("pki: PKI asset creation took %s", time.Since(s))
	}(start)
	pkiParams.ExtraFQDNs = append(pkiParams.ExtraFQDNs, "kubernetes")
	pkiParams.ExtraFQDNs = append(pkiParams.ExtraFQDNs, "kubernetes.default")
	pkiParams.ExtraFQDNs = append(pkiParams.ExtraFQDNs, "kubernetes.default.svc")
	pkiParams.ExtraFQDNs = append(pkiParams.ExtraFQDNs, fmt.Sprintf("kubernetes.default.svc.%s", pkiParams.ClusterDomain))
	pkiParams.ExtraFQDNs = append(pkiParams.ExtraFQDNs, "kubernetes.kube-system")
	pkiParams.ExtraFQDNs = append(pkiParams.ExtraFQDNs, "kubernetes.kube-system.svc")
	pkiParams.ExtraFQDNs = append(pkiParams.ExtraFQDNs, fmt.Sprintf("kubernetes.kube-system.svc.%s", pkiParams.ClusterDomain))

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
	caCertificate, err = pemToCertificate(pkiParams.CaPair.CertificatePem)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	caPrivateKey, err = pemToKey(pkiParams.CaPair.PrivateKeyPem)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	group.Go(func() (err error) {
		certPram := certParams{
			commonName:    "apiserver",
			caCertificate: caCertificate,
			caPrivateKey:  caPrivateKey,
			isEtcd:        false,
			isServer:      true,
			extraFQDNs:    pkiParams.ExtraFQDNs,
			extraIPs:      pkiParams.ExtraIPs,
			organization:  nil,
			keySize:       pkiParams.PkiKeySize,
		}
		apiServerCertificate, apiServerPrivateKey, err = createCertificate(certPram)
		return err
	})

	group.Go(func() (err error) {
		organization := make([]string, 1)
		organization[0] = "system:masters"
		certPram := certParams{
			commonName:    "client",
			caCertificate: caCertificate,
			caPrivateKey:  caPrivateKey,
			isEtcd:        false,
			isServer:      false,
			extraFQDNs:    nil,
			extraIPs:      nil,
			organization:  organization,
			keySize:       pkiParams.PkiKeySize,
		}
		clientCertificate, clientPrivateKey, err = createCertificate(certPram)
		return err
	})

	group.Go(func() (err error) {
		organization := make([]string, 1)
		organization[0] = "system:masters"

		certPram := certParams{
			commonName:    "client",
			caCertificate: caCertificate,
			caPrivateKey:  caPrivateKey,
			isEtcd:        false,
			isServer:      false,
			extraFQDNs:    nil,
			extraIPs:      nil,
			organization:  organization,
			keySize:       pkiParams.PkiKeySize,
		}

		kubeConfigCertificate, kubeConfigPrivateKey, err = createCertificate(certPram)
		return err
	})

	group.Go(func() (err error) {
		certPram := certParams{
			commonName:    "etcdserver",
			caCertificate: caCertificate,
			caPrivateKey:  caPrivateKey,
			isEtcd:        true,
			isServer:      true,
			extraFQDNs:    nil,
			extraIPs:      pkiParams.ExtraIPs,
			organization:  nil,
			keySize:       pkiParams.PkiKeySize,
		}
		etcdServerCertificate, etcdServerPrivateKey, err = createCertificate(certPram)
		return err
	})

	group.Go(func() (err error) {
		certPram := certParams{
			commonName:    "etcdclient",
			caCertificate: caCertificate,
			caPrivateKey:  caPrivateKey,
			isEtcd:        true,
			isServer:      false,
			extraFQDNs:    nil,
			extraIPs:      pkiParams.ExtraIPs,
			organization:  nil,
			keySize:       pkiParams.PkiKeySize,
		}
		etcdClientCertificate, etcdClientPrivateKey, err = createCertificate(certPram)
		return err
	})

	etcdPeerCertPairs = make([]*PkiKeyCertPair, pkiParams.MasterCount)
	for i := 0; i < pkiParams.MasterCount; i++ {
		i := i
		group.Go(func() (err error) {
			certPram := certParams{
				commonName:    "etcdpeer",
				caCertificate: caCertificate,
				caPrivateKey:  caPrivateKey,
				isEtcd:        true,
				isServer:      false,
				extraFQDNs:    nil,
				extraIPs:      pkiParams.ExtraIPs,
				organization:  nil,
				keySize:       pkiParams.PkiKeySize,
			}
			etcdPeerCertificate, etcdPeerPrivateKey, err := createCertificate(certPram)
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

type certParams struct {
	commonName    string
	caCertificate *x509.Certificate
	caPrivateKey  *rsa.PrivateKey
	isEtcd        bool
	isServer      bool
	extraFQDNs    []string
	extraIPs      []net.IP
	organization  []string
	keySize       int
}

func createCertificate(options certParams) (*x509.Certificate, *rsa.PrivateKey, error) {
	var err error

	isCA := (options.caCertificate == nil)

	now := time.Now()

	template := x509.Certificate{
		Subject:   pkix.Name{CommonName: options.commonName},
		NotBefore: now,
		NotAfter:  now.Add(ValidityDuration),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	if options.organization != nil {
		template.Subject.Organization = options.organization
	}

	if isCA {
		template.KeyUsage |= x509.KeyUsageCertSign
		template.IsCA = isCA
	} else if options.isEtcd {
		if options.commonName == "etcdServer" {
			template.IPAddresses = options.extraIPs
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
		} else if options.commonName == "etcdClient" {
			template.IPAddresses = options.extraIPs
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
		} else {
			template.IPAddresses = options.extraIPs
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
			template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
		}
	} else if options.isServer {
		template.DNSNames = options.extraFQDNs
		template.IPAddresses = options.extraIPs
		template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
	} else {
		template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
	}

	snMax := new(big.Int).Lsh(big.NewInt(1), 128)
	template.SerialNumber, err = rand.Int(rand.Reader, snMax)
	if err != nil {
		return nil, nil, err
	}

	privateKey, _ := rsa.GenerateKey(rand.Reader, options.keySize)

	var privateKeyToUse *rsa.PrivateKey
	var certificateToUse *x509.Certificate
	if !isCA {
		privateKeyToUse = options.caPrivateKey
		certificateToUse = options.caCertificate
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
	_ = pem.Encode(&pemBuffer, pemBlock)

	return pemBuffer.Bytes()
}

func privateKeyToPem(privateKey *rsa.PrivateKey) []byte {
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	pemBuffer := bytes.Buffer{}
	_ = pem.Encode(&pemBuffer, pemBlock)

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
