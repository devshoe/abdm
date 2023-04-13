package abdm

import (
	"fmt"

	"github.com/devshoe/abdm/pkg/abdm/models"
)

// Authenticate retrieves access token and refresh token if credentials are valid
func (c *Client) Authenticate() (err error) {
	var (
		request  models.SessionRequest
		response models.SessionResponse
		method                     = "POST"
		headers  map[string]string = nil
		code     int
	)

	request.ClientID = c.ClientID
	request.ClientSecret = c.ClientSecret

	if code, err = do(c.c, method, URISession, headers, request, &response); err != nil {
		return err
	} else if code != 200 {
		fmt.Println(response)
		return fmt.Errorf("authenticate: response code %d", code)
	}

	c.AccessToken = response.AccessToken
	c.RefreshToken = response.RefreshToken

	return
}
