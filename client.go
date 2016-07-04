package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/vault/api"
)

func getClient(vaultAddress string) *api.Client {
	c := api.Config{Address: vaultAddress}
	// response := c.ReadEnvironment()
	// if response != nil {
	// 	log.Error("Error reading local env: ", response.Error())
	// }
	client, err := api.NewClient(&c)
	if err != nil {
		log.Error("Error creating Vault Client: ", err.Error())
	}
	return client
}
