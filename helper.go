package vaultutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func readerToString(data io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	return buf.String()
}

func getClientParams(c *vaultClient) error {
	c.base_url = os.Getenv("VAULT_ADDR")
	if c.base_url == "" {
		return &argError{"VAULT_ADDR environment variables needs to be set"}
	}

	if c.credtype == "userpass" {
		c.username = os.Getenv("VAULT_USER")
		c.password = os.Getenv("VAULT_PASSWORD")
		if c.username == "" {
			return &argError{"VAULT_USER environment variables needs to be set"}
		}
		if c.password == "" {
			return &argError{"VAULT_PASSWORD environment variables needs to be set"}
		}
	} else if c.credtype == "ldap" {
		c.username = os.Getenv("VAULT_LUSER")
		c.password = os.Getenv("VAULT_LPASSWORD")
		if c.username == "" {
			return &argError{"VAULT_LUSER environment variables needs to be set"}
		}
		if c.password == "" {
			return &argError{"VAULT_LPASSWORD environment variables needs to be set"}
		}
	} else if c.credtype == "" {
		return &argError{"CRED type is not set"}
	}

	return nil
}

func GetVaultClient(credtype string) (*vaultClient, error) {
	client := &vaultClient{}
	client.credtype = credtype
	err := getClientParams(client)
	return client, err
}

func (c *vaultClient) GetToken() string {
	return c.token
}

func parseAuthResponse(resp *http.Response) (*authResponse, error) {
	defer resp.Body.Close()
	jsonResponse := new(authResponse)
	err := json.NewDecoder(resp.Body).Decode(&jsonResponse)
	return jsonResponse, err
}
