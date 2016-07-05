package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/vault/api"
)

type vaultClientInput struct {
	Address    string
	Token      string
	SSLEnabled bool
	SSLVerify  bool
	SSLCert    string
	SSLKey     string
	SSLCACert  string
}

func getClient(input *vaultClientInput) *api.Client {
	config := api.DefaultConfig()

	if input.Address != "" {
		log.Info("Setting Vault Address to: ", input.Address)
		config.Address = input.Address
	}

	transport := cleanhttp.DefaultTransport()

	if input.SSLEnabled {
		log.Info("Enabling SSL for Vault Client")
		var tlsConfig tls.Config

		if input.SSLCert != "" && input.SSLKey != "" {
			cert, err := tls.LoadX509KeyPair(input.SSLCert, input.SSLKey)
			if err != nil {
				log.Error("Error setting SSL Cert and Key")
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		} else if input.SSLCert != "" {
			cert, err := tls.LoadX509KeyPair(input.SSLCert, input.SSLCert)
			if err != nil {
				log.Error("Error setting SSL Cert")
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}

		if input.SSLCACert != "" {
			cacert, err := ioutil.ReadFile(input.SSLCACert)
			if err != nil {
				log.Error("Error setting SSL Cert CA")
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(cacert)
			tlsConfig.RootCAs = caCertPool
		}

		tlsConfig.BuildNameToCertificate()

		if !input.SSLVerify {
			log.Debug("Turning off SSL verification")
			tlsConfig.InsecureSkipVerify = true
		}

		transport.TLSClientConfig = &tlsConfig
	}

	config.HttpClient.Transport = transport

	client, err := api.NewClient(config)
	if err != nil {
		log.Error("Error creating Vault Client: ", err.Error())
	}

	if input.Token != "" {
		log.Info("Setting Client token")
		client.SetToken(input.Token)
	}

	return client
}
