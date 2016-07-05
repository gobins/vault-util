package main

import (
	log "github.com/Sirupsen/logrus"
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
	config.Address = input.Address

	client, err := api.NewClient(config)
	if err != nil {
		log.Error("Error creating Vault Client: ", err.Error())
	}
	return client
}
