package vaultutil

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func EnableAuthBackend(c *vaultapi.Client, mount_path, authType, desc string) error {
	err := c.Sys().EnableAuth(mount_path, authType, desc)
	if err != nil {
		return err
	}
	return err
}

func DisableAuthBackend(c *vaultapi.Client, mount_path string) error {
	err := c.Sys().DisableAuth(mount_path)
	if err != nil {
		return err
	}
	return err
}
