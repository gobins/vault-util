# vaultutil
A utility tool which uses the Haschicorp Vault API.
This tool can be used to authenticate using Vault's HTTP api.

## Support
vaultutil has been tested with the following versions
* 0.6

## Usage
go get github.com/gobins/vaultutil

### Creating a client
client, err := GetVaultClient("userpass")

### LDAP Authentication
* VAULT_LUSER
* VAULT_LPASSWORD

### Userpass Authentication
* VAULT_USER
* VAULT_PASSWORD

client.Authenticate()

### Retrieving Auth Token
token = client.GetToken()


## Features
* Authenticate using ldap backend
* Authenticate using userpass backend

## Contributing
Please contribute by sending a pull request.
If there is an issue due to API changes, please raise an issue.
