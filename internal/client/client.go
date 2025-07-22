package client

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
)

func InitGigaChatClient() *http.Client {
	certPool := x509.NewCertPool()
	pemData, _ := os.ReadFile("certs/giga.pem")
	crtData, _ := os.ReadFile("certs/giga.crt")
	certPool.AppendCertsFromPEM(pemData)
	certPool.AppendCertsFromPEM(crtData)

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
