package gofortiadc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"testing"
	"time"
)

func TestClient_SystemGetLocalCertificates(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.SystemGetLocalCertificates()
	if err == nil {
		t.Logf("%+v", res)
		t.Fatal(err)
	}
}

func TestClient_SystemGetLocalCertificate(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.SystemGetLocalCertificate("Factory")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SystemCreateLocalCertificate(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	cert, key, err := generateCertificate()
	if err != nil {
		t.Fatal(err)
	}

	err = client.SystemCreateLocalCertificate("gofortiadc", "", cert, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SystemDeleteLocalCertificate(t *testing.T) {
	if os.Getenv("TEST_LENS") != "true" {
		t.Skip()
	}

	client, err := NewClientHelper()
	if err != nil {
		t.Fatal(err)
	}

	err = client.SystemDeleteLocalCertificate("gofortiadc")
	if err != nil {
		t.Fatal(err)
	}
}

func generateCertificate() (cert, key []byte, err error) {

	rawKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return cert, key, err
	}

	template := x509.Certificate{
		Subject: pkix.Name{
			Organization: []string{"Gofortiadc Test"},
		},
		SerialNumber:          big.NewInt(1),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Minute * 30),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"gofortiadc.test.tld"},
	}

	rawCert, err := x509.CreateCertificate(rand.Reader, &template, &template, &rawKey.PublicKey, rawKey)
	if err != nil {
		return cert, key, err
	}

	cert = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: rawCert})
	key = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rawKey)})

	return cert, key, nil
}
