// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net"
	"testing"
)

func TestCreateCertificateWithOrganisation(t *testing.T) {
	var err error
	var caPair *PkiKeyCertPair

	var (
		caCertificate   *x509.Certificate
		caPrivateKey    *rsa.PrivateKey
		testCertificate *x509.Certificate
	)

	caCertificate, caPrivateKey, err = createCertificate("ca", nil, nil, false, false, nil, nil, nil)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}
	caPair = &PkiKeyCertPair{CertificatePem: string(certificateToPem(caCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(caPrivateKey))}

	caCertificate, err = pemToCertificate(caPair.CertificatePem)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}
	caPrivateKey, err = pemToKey(caPair.PrivateKeyPem)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}

	organization := make([]string, 1)
	organization[0] = "system:masters"
	testCertificate, _, err = createCertificate("client", caCertificate, caPrivateKey, false, false, nil, nil, organization)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}

	certificationOrganization := testCertificate.Subject.Organization

	if certificationOrganization[0] != organization[0] || len(certificationOrganization) != len(organization) {
		t.Fatalf("certificate organisation did not match")
	}
}

func TestCreateCertificateWithoutOrganisation(t *testing.T) {
	var err error
	var caPair *PkiKeyCertPair

	var (
		caCertificate   *x509.Certificate
		caPrivateKey    *rsa.PrivateKey
		testCertificate *x509.Certificate
	)

	caCertificate, caPrivateKey, err = createCertificate("ca", nil, nil, false, false, nil, nil, nil)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}
	caPair = &PkiKeyCertPair{CertificatePem: string(certificateToPem(caCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(caPrivateKey))}

	caCertificate, err = pemToCertificate(caPair.CertificatePem)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}
	caPrivateKey, err = pemToKey(caPair.PrivateKeyPem)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}

	testCertificate, _, err = createCertificate("client", caCertificate, caPrivateKey, false, false, nil, nil, nil)
	if err != nil {
		t.Fatalf("failed to generate certificate: %s", err)
	}

	certificationOrganization := testCertificate.Subject.Organization

	if len(certificationOrganization) != 0 {
		t.Fatalf("certificate organisation should be empty but has length %d", len(certificationOrganization))
	}
}

func TestSubjectAltNameInCert(t *testing.T) {
	extraIPs := []net.IP{net.ParseIP("10.255.255.5"), net.ParseIP("10.255.255.15")}
	roots := x509.NewCertPool()

	// Prepare CA and add it to certificate store.
	caCertificate, caPrivateKey, err := createCertificate("ca", nil, nil, false, false, nil, nil, nil)
	if err != nil {
		t.Fatalf("failed to generate CA certificates: %s.", err)
	}
	caPair := &PkiKeyCertPair{CertificatePem: string(certificateToPem(caCertificate.Raw)), PrivateKeyPem: string(privateKeyToPem(caPrivateKey))}

	ok := roots.AppendCertsFromPEM([]byte(caPair.CertificatePem))
	if !ok {
		t.Fatalf("failed to parse generated CA certificate.")
	}

	// Test SAN is not empty
	SubjectAltNames := []string{"santest.mydomain.com", "santest2.testdomain.net"}
	formattedDNSPrefixes := []string{
		"santest.australiacentral.cloudapp.azure.com",
		"santest.australiacentral2.cloudapp.azure.com",
		"santest.australiaeast.cloudapp.azure.com",
		"santest.australiasoutheast.cloudapp.azure.com",
		"santest.brazilsouth.cloudapp.azure.com",
		"santest.canadacentral.cloudapp.azure.com",
		"santest.canadaeast.cloudapp.azure.com",
		"santest.centralindia.cloudapp.azure.com",
		"santest.centralus.cloudapp.azure.com",
		"santest.centraluseuap.cloudapp.azure.com",
		"santest.chinaeast.cloudapp.chinacloudapi.cn",
		"santest.chinaeast2.cloudapp.chinacloudapi.cn",
		"santest.chinanorth.cloudapp.chinacloudapi.cn",
		"santest.chinanorth2.cloudapp.chinacloudapi.cn",
		"santest.eastasia.cloudapp.azure.com",
		"santest.eastus.cloudapp.azure.com",
		"santest.eastus2.cloudapp.azure.com",
		"santest.eastus2euap.cloudapp.azure.com",
		"santest.francecentral.cloudapp.azure.com",
		"santest.francesouth.cloudapp.azure.com",
		"santest.japaneast.cloudapp.azure.com",
		"santest.japanwest.cloudapp.azure.com",
		"santest.koreacentral.cloudapp.azure.com",
		"santest.koreasouth.cloudapp.azure.com",
		"santest.northcentralus.cloudapp.azure.com",
		"santest.northeurope.cloudapp.azure.com",
		"santest.southcentralus.cloudapp.azure.com",
		"santest.southeastasia.cloudapp.azure.com",
		"santest.southindia.cloudapp.azure.com",
		"santest.uksouth.cloudapp.azure.com",
		"santest.ukwest.cloudapp.azure.com",
		"santest.westcentralus.cloudapp.azure.com",
		"santest.westeurope.cloudapp.azure.com",
		"santest.westindia.cloudapp.azure.com",
		"santest.westus.cloudapp.azure.com",
		"santest.westus2.cloudapp.azure.com",
		"santest.chinaeast.cloudapp.chinacloudapi.cn",
		"santest.chinanorth.cloudapp.chinacloudapi.cn",
		"santest.chinanorth2.cloudapp.chinacloudapi.cn",
		"santest.chinaeast2.cloudapp.chinacloudapi.cn",
		"santest.germanycentral.cloudapp.microsoftazure.de",
		"santest.germanynortheast.cloudapp.microsoftazure.de",
		"santest.usgovvirginia.cloudapp.usgovcloudapi.net",
		"santest.usgoviowa.cloudapp.usgovcloudapi.net",
		"santest.usgovarizona.cloudapp.usgovcloudapi.net",
		"santest.usgovtexas.cloudapp.usgovcloudapi.net",
		"santest.francecentral.cloudapp.azure.com",
	}
	masterExtraFQDNs := append(formattedDNSPrefixes, SubjectAltNames...)
	expectedDNSNames := []string{"santest.australiaeast.cloudapp.azure.com", "santest.australiasoutheast.cloudapp.azure.com", "santest.brazilsouth.cloudapp.azure.com", "santest.canadacentral.cloudapp.azure.com", "santest.canadaeast.cloudapp.azure.com", "santest.centralindia.cloudapp.azure.com", "santest.centralus.cloudapp.azure.com", "santest.centraluseuap.cloudapp.azure.com", "santest.chinaeast.cloudapp.chinacloudapi.cn", "santest.chinaeast2.cloudapp.chinacloudapi.cn", "santest.chinanorth.cloudapp.chinacloudapi.cn", "santest.chinanorth2.cloudapp.chinacloudapi.cn", "santest.eastasia.cloudapp.azure.com", "santest.eastus.cloudapp.azure.com", "santest.eastus2.cloudapp.azure.com", "santest.eastus2euap.cloudapp.azure.com", "santest.japaneast.cloudapp.azure.com", "santest.japanwest.cloudapp.azure.com", "santest.koreacentral.cloudapp.azure.com", "santest.koreasouth.cloudapp.azure.com", "santest.northcentralus.cloudapp.azure.com", "santest.northeurope.cloudapp.azure.com", "santest.southcentralus.cloudapp.azure.com", "santest.southeastasia.cloudapp.azure.com", "santest.southindia.cloudapp.azure.com", "santest.uksouth.cloudapp.azure.com", "santest.ukwest.cloudapp.azure.com", "santest.westcentralus.cloudapp.azure.com", "santest.westeurope.cloudapp.azure.com", "santest.westindia.cloudapp.azure.com", "santest.westus.cloudapp.azure.com", "santest.westus2.cloudapp.azure.com", "santest.chinaeast.cloudapp.chinacloudapi.cn", "santest.chinanorth.cloudapp.chinacloudapi.cn", "santest.germanycentral.cloudapp.microsoftazure.de", "santest.germanynortheast.cloudapp.microsoftazure.de", "santest.usgovvirginia.cloudapp.usgovcloudapi.net", "santest.usgoviowa.cloudapp.usgovcloudapi.net", "santest.usgovarizona.cloudapp.usgovcloudapi.net", "santest.usgovtexas.cloudapp.usgovcloudapi.net", "santest.francecentral.cloudapp.azure.com", "santest.mydomain.com", "santest2.testdomain.net", "kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster.local", "kubernetes.kube-system", "kubernetes.kube-system.svc", "kubernetes.kube-system.svc.cluster.local"}

	apiServerPair, _, _, _, _, _, err := CreatePki(masterExtraFQDNs, extraIPs, "cluster.local", caPair, 1)

	if err != nil {
		t.Fatalf("failed to generate certificates: %s.", err)
	}

	if apiServerPair != nil {
		block, _ := pem.Decode([]byte(apiServerPair.CertificatePem))
		if block == nil {
			panic("failed to parse certificate PEM")
		}
		var cert *x509.Certificate
		cert, err = x509.ParseCertificate(block.Bytes)

		if err != nil {
			t.Fatalf("Can not parse generated certificate %s", err)
		}

		for _, domain := range expectedDNSNames {
			opts := x509.VerifyOptions{
				DNSName: domain,
				Roots:   roots,
			}

			if _, err = cert.Verify(opts); err != nil {
				t.Fatalf("failed to verify certificate: %s", err)
			}
		}
	} else {
		t.Fatalf("API server pair not generated.")
	}

	// Test SAN is empty
	SubjectAltNames = []string{}
	masterExtraFQDNs = append(formattedDNSPrefixes, SubjectAltNames...)
	expectedDNSNames = []string{"santest.australiaeast.cloudapp.azure.com", "santest.australiasoutheast.cloudapp.azure.com", "santest.brazilsouth.cloudapp.azure.com", "santest.canadacentral.cloudapp.azure.com", "santest.canadaeast.cloudapp.azure.com", "santest.centralindia.cloudapp.azure.com", "santest.centralus.cloudapp.azure.com", "santest.centraluseuap.cloudapp.azure.com", "santest.chinaeast.cloudapp.chinacloudapi.cn", "santest.chinaeast2.cloudapp.chinacloudapi.cn", "santest.chinanorth.cloudapp.chinacloudapi.cn", "santest.chinanorth2.cloudapp.chinacloudapi.cn", "santest.eastasia.cloudapp.azure.com", "santest.eastus.cloudapp.azure.com", "santest.eastus2.cloudapp.azure.com", "santest.eastus2euap.cloudapp.azure.com", "santest.japaneast.cloudapp.azure.com", "santest.japanwest.cloudapp.azure.com", "santest.koreacentral.cloudapp.azure.com", "santest.koreasouth.cloudapp.azure.com", "santest.northcentralus.cloudapp.azure.com", "santest.northeurope.cloudapp.azure.com", "santest.southcentralus.cloudapp.azure.com", "santest.southeastasia.cloudapp.azure.com", "santest.southindia.cloudapp.azure.com", "santest.uksouth.cloudapp.azure.com", "santest.ukwest.cloudapp.azure.com", "santest.westcentralus.cloudapp.azure.com", "santest.westeurope.cloudapp.azure.com", "santest.westindia.cloudapp.azure.com", "santest.westus.cloudapp.azure.com", "santest.westus2.cloudapp.azure.com", "santest.chinaeast.cloudapp.chinacloudapi.cn", "santest.chinanorth.cloudapp.chinacloudapi.cn", "santest.germanycentral.cloudapp.microsoftazure.de", "santest.germanynortheast.cloudapp.microsoftazure.de", "santest.usgovvirginia.cloudapp.usgovcloudapi.net", "santest.usgoviowa.cloudapp.usgovcloudapi.net", "santest.usgovarizona.cloudapp.usgovcloudapi.net", "santest.usgovtexas.cloudapp.usgovcloudapi.net", "santest.francecentral.cloudapp.azure.com", "kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster.local", "kubernetes.kube-system", "kubernetes.kube-system.svc", "kubernetes.kube-system.svc.cluster.local"}

	apiServerPair, _, _, _, _, _, err = CreatePki(masterExtraFQDNs, extraIPs, "cluster.local", caPair, 1)

	if err != nil {
		t.Fatalf("failed to generate certificates: %s.", err)
	}

	if apiServerPair != nil {
		block, _ := pem.Decode([]byte(apiServerPair.CertificatePem))
		if block == nil {
			panic("failed to parse certificate PEM")
		}
		cert, err := x509.ParseCertificate(block.Bytes)

		if err != nil {
			t.Fatalf("Can not parse generated certificate %s", err)
		}

		for _, domain := range expectedDNSNames {
			opts := x509.VerifyOptions{
				DNSName: domain,
				Roots:   roots,
			}

			if _, err := cert.Verify(opts); err != nil {
				t.Fatalf("failed to verify certificate: %s", err)
			}
		}
	} else {
		t.Fatalf("API server pair not generated.")
	}
}

func TestCreatePkiKeyCertPair(t *testing.T) {
	subject := "foosubject"
	_, err := CreatePkiKeyCertPair(subject)
	if err != nil {
		t.Errorf("unexpected error thrown while executing CreatePkiKeyCertPair : %s", err.Error())
	}
}
