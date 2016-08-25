# vaultutil
A utility tool which uses the Haschicorp Vault API.
This tool can be used to authenticate using Vault's HTTP api.

## Features
* Authenticate using ldap backend
* Authenticate using userpass backend
* HTTP Api call support for custom paths


## Usage
go get github.com/gobins/vaultutil

The following environment variables needs to be set for authentication
### LDAP Authentication
* VAULT_LUSER
* VAULT_LPASSWORD

### Userpass Authentication
* VAULT_USER
* VAULT_PASSWORD


### Creating a client
client, err := GetVaultClient("userpass")

or

client, err := GetVaultClient("ldap")

client.Authenticate()

### Retrieving Auth Token
token = client.GetToken()

## Support
vaultutil has been tested with the following versions
* 0.6
* 0.6.1


## Contributing
Please contribute by sending a pull request.

If there is an issue due to API changes, please raise an issue.
