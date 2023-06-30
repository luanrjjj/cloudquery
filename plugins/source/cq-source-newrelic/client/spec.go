package client

import "errors"

type Spec struct {
	// plugin spec goes here
	Accounts []Account `json:"accounts"`
	Token    string    `json:"access_token,omitempty"`
}

type Account struct {
	Name   string `json:"name"`
	ApiKey string `json:"api_key"`
	AppKey string `json:"app_key"`
}

func (s Spec) Validate() error {
	newrelicToken := s.Token
	if newrelicToken == "" {
		return errors.New("missing GitLab API token in configuration file")
	}
	return nil
}
