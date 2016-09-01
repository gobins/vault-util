package vaultutil

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func CreateMount(c *vaultapi.Client, mount_path, mount_type string) error {
	data := map[string]interface{}{"type": mount_type}

	_, err := c.Logical().Write("sys/mounts/"+mount_path, data)
	if err != nil {
		return err
	}
	return nil
}

func DoRemount(c *vaultapi.Client, mount_path, new_mount_path string) error {
	data := map[string]interface{}{"from": mount_path, "to": new_mount_path}

	_, err := c.Logical().Write("sys/remount", data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMount(c *vaultapi.Client, mount_path string) error {
	_, err := c.Logical().Delete("sys/mounts/" + mount_path)
	if err != nil {
		return err
	}
	return nil
}

func ListMounts(c *vaultapi.Client) (map[string]*api.MountOutput, error) {
	data, err := c.Sys().ListMounts()
	if err != nil {
		return data, err
	}
	return data, err
}
