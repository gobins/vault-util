package vaultutil

const VaultAddress = "VAULT_ADDR"

type authResponse struct {
	Lease_Id       string `json:"lease_id"`
	Renewable      bool   `json:"renewable"`
	Lease_Duration int    `json:"lease_duration"`
	//Data           string `json:"data"`
	Wrap_Info  string `json:"wrap_info"`
	Warnings   string `json:"warnings"`
	AuthHolder `json:"auth"`
}

type SSHRoleData struct {
	Key               string
	Admin_User        string
	Default_User      string
	Cidr_List         string
	Exclude_Cidr_List string
	Port              int
	key_type          string
	key_bits          int
	install_script    string
	allowed_users     string
	key_option_specs  string
}

type AuthHolder struct {
	Client_Token   string   `json:"client_token"`
	Accessor       string   `json:"accessor"`
	Policies       []string `json:"policies"`
	Metadata       `json:"metadata"`
	Lease_Duration int  `json:"lease_duration"`
	Renewable      bool `json:"renewable"`
}

type Metadata struct {
	Username string `json:"username"`
	Policies string `json:"policies"`
}

type vaultClient struct {
	base_url  string
	uri       string
	sslEnable bool
	sslVerify bool
	data      string
	username  string
	password  string
	credtype  string
	token     string
}

type argError struct {
	s string
}

func (e *argError) Error() string {
	return e.s
}

type authError struct {
	s string
}

func (e *authError) Error() string {
	return e.s
}
