package vaultutil

import (
	"crypto/tls"
	"net/http"
	"strings"
)

func (c *vaultClient) Get() (*http.Response, error) {

	tr := &http.Transport{}
	url := c.base_url + c.uri
	if c.sslEnable && !c.sslVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	vclient := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)

	if c.token != "" {

		req.Header.Set("X-Vault-Token", c.token)
	}
	resp, err := vclient.Do(req)
	return resp, err
}

func (c *vaultClient) Post() (*http.Response, error) {
	tr := &http.Transport{}

	if c.sslEnable && !c.sslVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	vclient := &http.Client{Transport: tr}
	url := c.base_url + c.uri
	req, _ := http.NewRequest("POST", url, strings.NewReader(c.data))
	if c.token != "" {

		req.Header.Set("X-Vault-Token", c.token)
	}

	resp, err := vclient.Do(req)

	return resp, err

}

func (c *vaultClient) Authenticate() (*authResponse, error) {

	if c.credtype == "userpass" {
		c.uri = "/v1/auth/userpass/login/" + c.username
		c.data = `{"password" : "` + c.password + `"}`
	}
	if c.credtype == "ldap" {
		c.uri = "/v1/auth/ldap/login/" + c.username
		c.data = `{"password" : "` + c.password + `"}`
	}

	resp, err := c.Post()
	jsonResponse := &authResponse{}

	if resp.StatusCode != 200 {
		return jsonResponse, &authError{readerToString(resp.Body)}
	}
	jsonResponse, err = parseAuthResponse(resp)
	c.token = jsonResponse.Client_Token
	return jsonResponse, err

}
