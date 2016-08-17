package vaultutil

import (
	"os"
	"testing"
)

func TestGetVaultClient(t *testing.T) {
	os.Setenv(VaultAddress, "https://vault.mycompany.com")
	client, err := GetVaultClient("ldap")
	if client.base_url != os.Getenv(VaultAddress) {
		t.Error("Client did not get Vault Server Address")
	}

	if err != nil {
		t.Error(err.Error())
	}
}
