package vaultutil

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func GetGroups(c *vaultapi.Client) (*api.Secret, error) {
	resp, err := c.Logical().List("auth/ldap/groups")
	if err != nil {
		return resp, err
	}
	return resp, err
}
