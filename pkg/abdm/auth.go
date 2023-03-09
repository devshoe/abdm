package abdm

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AuthenticationRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type AuthenticationResponse struct {
	AccessToken      string `json:"accessToken"`
	ExpiresIn        int    `json:"expiresIn"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresIn int    `json:"refreshExpiresIn"`
	TokenType        string `json:"tokenType"`
}

func (c *Client) Authenticate() error {
	var (
		requestbody  AuthenticationRequest
		responsebody AuthenticationResponse
		url          = "https://dev.abdm.gov.in/gateway/v0.5/sessions"
	)
	requestbody.ClientID = c.ClientID
	requestbody.ClientSecret = c.ClientSecret

	payloadBytes, _ := json.Marshal(requestbody)

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	request.Header.Set("Content-Type", "application/json")
	if response, err := c.c.Do(request); err != nil {
		return err
	} else {
		defer response.Body.Close()

		if err := json.NewDecoder(response.Body).Decode(&responsebody); err != nil {
			return err
		}
		c.AccessToken = responsebody.AccessToken
		c.RefreshToken = responsebody.RefreshToken
	}

	return nil

}
