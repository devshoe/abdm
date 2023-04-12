package abdm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	c *http.Client

	ClientID     string
	ClientSecret string

	AccessToken  string
	RefreshToken string

	HIURegistrationID string
	HIPRegistrationID string
	ConsentManagerID  string // sbx for sandbox, ndhm for production

}

func New(clientID, clientSecret string) *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		c:            &http.Client{Jar: jar},
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (c *Client) do(method, url string, headers map[string]string, requestBody, responseBody any) error {
	var request *http.Request

	if requestBody != nil {
		payloadBytes, _ := json.Marshal(requestBody)
		request, _ = http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	} else {
		request, _ = http.NewRequest(method, url, nil)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.AccessToken)

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	if response, err := c.c.Do(request); err != nil {
		return err
	} else {
		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(responseBody); err != nil {
			return err
		}
	}

	return nil
}
