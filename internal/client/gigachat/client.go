package gigachat

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
)

type GigaChat struct {
	httpClient *http.Client
}

func NewGigaChat() *GigaChat {
	return &GigaChat{httpClient: newHttpClient()}
}

func newHttpClient() *http.Client {
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
