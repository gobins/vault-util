package vaultutil

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func CreateUpdateSSHKey(c *vaultapi.Client, ssh_path, key_name, key_value string) (*vaultapi.Secret, error) {
	data := map[string]interface{}{"key": key_value}
	secret, err := c.Logical().Write(ssh_path+key_name, data)
	return secret, err
}

func DeleteSSHKey(c *vaultapi.Client, ssh_path, key_name string) (*vaultapi.Secret, error) {
	secret, err := c.Logical().Delete(ssh_path + key_name)
	return secret, err
}

func CreateUpdateSSHRole(c *vaultapi.Client, ssh_path, role_name string, role SSHRoleData) (*vaultapi.Secret, error) {
	data := map[string]interface{}{
		"key":        role.Key,
		"admin_user": role.Admin_User,
		"defaul_user", role.Default_User,
		"cidr_list ":        role.Cidr_List,
		"exclude_cidr_list": role.Exclude_Cidr_List,
		"port":              role.Port,
		"key_type":          role.key_type,
		"key_bits":          role.key_bits,
		"install_script":    role.install_script,
		"allowed_users":     role.allowed_users,
		"key_option_specs":  role.key_option_specs}
	secret, err := c.Logical().Write(ssh_path+key_name, data)
	return secret, err
}

func GetSSHRole(c *vaultapi.Client, ssh_path, role_name string) (*vaultapi.Secret, error) {
	secret, err := c.Logical().Read(ssh_path+key_name, data)
	return secret, err
}
