package abdm

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/devshoe/abdm/pkg/abdm/models"
)

func (c *Client) Authenticate() error {
	var (
		requestbody  models.AuthenticationRequest
		responsebody models.AuthenticationResponse
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
