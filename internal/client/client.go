package client

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
)

func InitGigaChatClient() *http.Client {
	certPool := x509.NewCertPool()
	firstCert, _ := os.ReadFile("/etc/whats-in-it/certs/first.pem")   // auth
	secondCert, _ := os.ReadFile("/etc/whats-in-it/certs/second.pem") // API
	certPool.AppendCertsFromPEM(firstCert)
	certPool.AppendCertsFromPEM(secondCert)

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := &http.Client{
		Transport: transport,
	}
	return client
}
