package vaultutil

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func GetGroups(c *vaultapi.Client) (*vaultapi.Secret, error) {
	resp, err := c.Logical().List("auth/ldap/groups")
	if err != nil {
		return resp, err
	}
	return resp, err
}

func GetGroupPolicy(c *vaultapi.Client, group_name string) (*vaultapi.Secret, error) {
	resp, err := c.Logical().Read("auth/ldap/groups/" + group_name)
	if err != nil {
		return resp, err
	}
	return resp, err
}
