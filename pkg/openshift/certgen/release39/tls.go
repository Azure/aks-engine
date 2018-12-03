// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package release39

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"time"

	"github.com/Azure/aks-engine/pkg/openshift/filesystem"
)

type authKeyID struct {
	KeyIdentifier             []byte      `asn1:"optional,tag:0"`
	AuthorityCertIssuer       generalName `asn1:"optional,tag:1"`
	AuthorityCertSerialNumber *big.Int    `asn1:"optional,tag:2"`
}

type generalName struct {
	DirectoryName pkix.RDNSequence `asn1:"optional,explicit,tag:4"`
}

func newCertAndKey(filename string, template, signingcert *x509.Certificate, signingkey *rsa.PrivateKey, etcdcaspecial, etcdclientspecial bool) (CertAndKey, error) {
	bits := 2048
	if etcdcaspecial {
		bits = 4096
	}

	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return CertAndKey{}, err
	}

	if signingcert == nil {
		// make it self-signed
		signingcert = template
		signingkey = key
	}

	if etcdcaspecial {
		template.SubjectKeyId = intsha1(key.N)
		ext := pkix.Extension{
			Id: []int{2, 5, 29, 35},
		}
		var err error
		ext.Value, err = asn1.Marshal(authKeyID{
			AuthorityCertIssuer:       generalName{DirectoryName: signingcert.Subject.ToRDNSequence()},
			AuthorityCertSerialNumber: signingcert.SerialNumber,
		})
		if err != nil {
			return CertAndKey{}, err
		}
		template.ExtraExtensions = append(template.Extensions, ext)
		template.MaxPathLenZero = true
	}

	if etcdclientspecial {
		template.SubjectKeyId = intsha1(key.N)
		ext := pkix.Extension{
			Id: []int{2, 5, 29, 35},
		}
		var err error
		ext.Value, err = asn1.Marshal(authKeyID{
			KeyIdentifier:             intsha1(signingkey.N),
			AuthorityCertIssuer:       generalName{DirectoryName: signingcert.Subject.ToRDNSequence()},
			AuthorityCertSerialNumber: signingcert.SerialNumber,
		})
		if err != nil {
			return CertAndKey{}, err
		}
		template.ExtraExtensions = append(template.Extensions, ext)
	}

	b, err := x509.CreateCertificate(rand.Reader, template, signingcert, key.Public(), signingkey)
	if err != nil {
		return CertAndKey{}, err
	}

	cert, err := x509.ParseCertificate(b)
	if err != nil {
		return CertAndKey{}, err
	}

	return CertAndKey{cert: cert, key: key}, nil
}

func certAsBytes(cert *x509.Certificate) ([]byte, error) {
	buf := &bytes.Buffer{}

	err := pem.Encode(buf, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func writeCert(fs filesystem.Writer, filename string, cert *x509.Certificate) error {
	b, err := certAsBytes(cert)
	if err != nil {
		return err
	}
	fi := GetFileInfo(filename)
	return fs.WriteFile(filename, b, fi)
}

func privateKeyAsBytes(key *rsa.PrivateKey) ([]byte, error) {
	buf := &bytes.Buffer{}

	err := pem.Encode(buf, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func writePrivateKey(fs filesystem.Writer, filename string, key *rsa.PrivateKey) error {
	b, err := privateKeyAsBytes(key)
	if err != nil {
		return err
	}

	fi := GetFileInfo(filename)
	return fs.WriteFile(filename, b, fi)
}

func writePublicKey(fs filesystem.Writer, filename string, key *rsa.PublicKey) error {
	buf := &bytes.Buffer{}

	b, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return err
	}

	err = pem.Encode(buf, &pem.Block{Type: "PUBLIC KEY", Bytes: b})
	if err != nil {
		return err
	}
	fi := GetFileInfo(filename)
	return fs.WriteFile(filename, buf.Bytes(), fi)
}

// PrepareMasterCerts creates the master certs
func (c *Config) PrepareMasterCerts() error {
	if c.cas == nil {
		c.cas = map[string]CertAndKey{}
	}

	if c.Master.certs == nil {
		c.Master.certs = map[string]CertAndKey{}
	}

	if c.Master.etcdcerts == nil {
		c.Master.etcdcerts = map[string]CertAndKey{}
	}

	ips := append([]net.IP{}, c.Master.IPs...)
	ips = append(ips, net.ParseIP("172.30.0.1"))

	dns := []string{
		c.ExternalMasterHostname, "kubernetes", "kubernetes.default", "kubernetes.default.svc",
		"kubernetes.default.svc.cluster.local", c.Master.Hostname, "openshift",
		"openshift.default", "openshift.default.svc",
		"openshift.default.svc.cluster.local",
	}
	for _, ip := range ips {
		dns = append(dns, ip.String())
	}

	now := time.Now()

	cacerts := []struct {
		filename string
		template *x509.Certificate
	}{
		{
			filename: "etc/origin/master/ca",
			template: &x509.Certificate{
				Subject: pkix.Name{CommonName: fmt.Sprintf("openshift-signer@%d", now.Unix())},
			},
		},
		{
			filename: "etc/origin/master/front-proxy-ca",
			template: &x509.Certificate{
				Subject: pkix.Name{CommonName: fmt.Sprintf("openshift-signer@%d", now.Unix())},
			},
		},
		{
			filename: "etc/origin/master/frontproxy-ca",
			template: &x509.Certificate{
				Subject: pkix.Name{CommonName: fmt.Sprintf("aggregator-proxy-car@%d", now.Unix())},
			},
		},
		{
			filename: "etc/origin/master/master.etcd-ca",
			template: &x509.Certificate{
				Subject: pkix.Name{CommonName: fmt.Sprintf("etcd-signer@%d", now.Unix())},
			},
		},
		{
			filename: "etc/origin/master/service-signer",
			template: &x509.Certificate{
				Subject: pkix.Name{CommonName: fmt.Sprintf("openshift-service-serving-signer@%d", now.Unix())},
			},
		},
		{
			filename: "etc/origin/service-catalog/ca",
			template: &x509.Certificate{
				Subject: pkix.Name{CommonName: "service-catalog-signer"},
			},
		},
	}

	for _, cacert := range cacerts {
		template := &x509.Certificate{
			SerialNumber:          c.serial.Get(),
			NotBefore:             now,
			NotAfter:              now.AddDate(5, 0, 0),
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign,
			BasicConstraintsValid: true,
			IsCA:                  true,
		}
		template.Subject = cacert.template.Subject

		certAndKey, err := newCertAndKey(cacert.filename, template, nil, nil, cacert.filename == "etc/origin/master/master.etcd-ca", false)
		if err != nil {
			return err
		}

		c.cas[cacert.filename] = certAndKey
	}

	certs := []struct {
		filename string
		template *x509.Certificate
		signer   string
	}{
		{
			filename: "etc/origin/master/admin",
			template: &x509.Certificate{
				Subject:     pkix.Name{Organization: []string{"system:cluster-admins", "system:masters"}, CommonName: "system:admin"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			},
		},
		{
			filename: "etc/origin/master/aggregator-front-proxy",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: "aggregator-front-proxy"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			},
			signer: "etc/origin/master/front-proxy-ca",
		},
		{
			filename: "etc/origin/master/etcd.server",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: c.Master.IPs[0].String()},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
				DNSNames:    dns,
				IPAddresses: ips,
			},
		},
		{
			filename: "etc/origin/master/master.etcd-client",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: c.Master.Hostname},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
				DNSNames:    []string{c.Master.Hostname}, // TODO
				IPAddresses: []net.IP{c.Master.IPs[0]},   // TODO
			},
			signer: "etc/origin/master/master.etcd-ca",
		},
		{
			filename: "etc/origin/master/master.kubelet-client",
			template: &x509.Certificate{
				Subject:     pkix.Name{Organization: []string{"system:node-admins"}, CommonName: "system:openshift-node-admin"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			},
		},
		{
			filename: "etc/origin/master/master.proxy-client",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: "system:master-proxy"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			},
		},
		{
			filename: "etc/origin/master/master.server",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: c.Master.IPs[0].String()},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
				DNSNames:    dns,
				IPAddresses: ips,
			},
		},
		{
			filename: "etc/origin/master/openshift-aggregator",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: "system:openshift-aggregator"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			},
			signer: "etc/origin/master/frontproxy-ca",
		},
		{
			filename: "etc/origin/master/openshift-master",
			template: &x509.Certificate{
				Subject:     pkix.Name{Organization: []string{"system:masters", "system:openshift-master"}, CommonName: "system:openshift-master"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			},
		},
		{
			filename: "etc/origin/master/node-bootstrapper",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: "system:serviceaccount:openshift-infra:node-bootstrapper"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			},
		},
		{
			filename: "etc/origin/service-catalog/apiserver",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: "apiserver.kube-service-catalog"},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
				DNSNames:    []string{"apiserver.kube-service-catalog", "apiserver.kube-service-catalog.svc", "apiserver.kube-service-catalog.svc.cluster.local"},
			},
			signer: "etc/origin/service-catalog/ca",
		},
		// TODO: registry cert
	}

	for _, cert := range certs {
		template := &x509.Certificate{
			SerialNumber:          c.serial.Get(),
			NotBefore:             now,
			NotAfter:              now.AddDate(2, 0, 0),
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
			BasicConstraintsValid: true,
		}
		template.Subject = cert.template.Subject
		template.ExtKeyUsage = cert.template.ExtKeyUsage
		template.DNSNames = cert.template.DNSNames
		template.IPAddresses = cert.template.IPAddresses

		if cert.signer == "" {
			cert.signer = "etc/origin/master/ca"
		}

		certAndKey, err := newCertAndKey(cert.filename, template, c.cas[cert.signer].cert, c.cas[cert.signer].key, false, cert.filename == "etc/origin/master/master.etcd-client")
		if err != nil {
			return err
		}

		c.Master.certs[cert.filename] = certAndKey
	}

	etcdcerts := []struct {
		filename string
		template *x509.Certificate
		signer   string
	}{
		{
			filename: "etc/etcd/peer",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: c.Master.Hostname},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
				DNSNames:    []string{c.Master.Hostname}, // TODO
				IPAddresses: []net.IP{c.Master.IPs[0]},   // TODO
			},
			signer: "etc/origin/master/master.etcd-ca",
		},
		{
			filename: "etc/etcd/server",
			template: &x509.Certificate{
				Subject:     pkix.Name{CommonName: c.Master.Hostname},
				ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
				DNSNames:    []string{c.Master.Hostname}, // TODO
				IPAddresses: []net.IP{c.Master.IPs[0]},   // TODO
			},
			signer: "etc/origin/master/master.etcd-ca",
		},
	}

	for _, cert := range etcdcerts {
		template := &x509.Certificate{
			SerialNumber:          c.serial.Get(),
			NotBefore:             now,
			NotAfter:              now.AddDate(5, 0, 0),
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
			BasicConstraintsValid: true,
		}
		template.Subject = cert.template.Subject
		template.ExtKeyUsage = cert.template.ExtKeyUsage
		template.DNSNames = cert.template.DNSNames
		template.IPAddresses = cert.template.IPAddresses

		certAndKey, err := newCertAndKey(cert.filename, template, c.cas[cert.signer].cert, c.cas[cert.signer].key, false, true)
		if err != nil {
			return err
		}

		c.Master.etcdcerts[cert.filename] = certAndKey
	}

	return nil
}

// WriteMasterCerts writes the master certs
func (c *Config) WriteMasterCerts(fs filesystem.Writer) error {
	for filename, ca := range c.cas {
		err := writeCert(fs, fmt.Sprintf("%s.crt", filename), ca.cert)
		if err != nil {
			return err
		}

		err = writePrivateKey(fs, fmt.Sprintf("%s.key", filename), ca.key)
		if err != nil {
			return err
		}
	}

	err := writeCert(fs, "etc/origin/master/ca-bundle.crt", c.cas["etc/origin/master/ca"].cert)
	if err != nil {
		return err
	}

	err = writeCert(fs, "etc/origin/master/client-ca-bundle.crt", c.cas["etc/origin/master/ca"].cert) // TODO: confirm if needed
	if err != nil {
		return err
	}

	err = writeCert(fs, "etc/etcd/ca.crt", c.cas["etc/origin/master/master.etcd-ca"].cert)
	if err != nil {
		return err
	}

	for filename, cert := range c.Master.certs {
		err := writeCert(fs, fmt.Sprintf("%s.crt", filename), cert.cert)
		if err != nil {
			return err
		}

		err = writePrivateKey(fs, fmt.Sprintf("%s.key", filename), cert.key)
		if err != nil {
			return err
		}
	}

	for filename, cert := range c.Master.etcdcerts {
		err := writeCert(fs, fmt.Sprintf("%s.crt", filename), cert.cert)
		if err != nil {
			return err
		}

		err = writePrivateKey(fs, fmt.Sprintf("%s.key", filename), cert.key)
		if err != nil {
			return err
		}
	}
	fname := "etc/origin/master/ca.serial.txt"
	fi := GetFileInfo(fname)
	return fs.WriteFile(fname, []byte(fmt.Sprintf("%02X\n", c.serial.Get())), fi)
}

// WriteBootstrapCerts writes the node bootstrap certs
func (c *Config) WriteBootstrapCerts(fs filesystem.Writer) error {
	err := writeCert(fs, "etc/origin/node/ca.crt", c.cas["etc/origin/master/ca"].cert)
	if err != nil {
		return err
	}

	err = writeCert(fs, "etc/origin/node/node-bootstrapper.crt", c.Master.certs["etc/origin/master/node-bootstrapper"].cert)
	if err != nil {
		return err
	}

	return writePrivateKey(fs, "etc/origin/node/node-bootstrapper.key", c.Master.certs["etc/origin/master/node-bootstrapper"].key)
}

// WriteMasterKeypair writes the master service account keypair
func (c *Config) WriteMasterKeypair(fs filesystem.Writer) error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	err = writePrivateKey(fs, "etc/origin/master/serviceaccounts.private.key", key)
	if err != nil {
		return err
	}

	return writePublicKey(fs, "etc/origin/master/serviceaccounts.public.key", &key.PublicKey)
}

func intsha1(n *big.Int) []byte {
	h := sha1.New()
	h.Write(n.Bytes())
	return h.Sum(nil)
}
