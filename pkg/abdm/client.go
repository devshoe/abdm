package abdm

import (
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

func (c *Client) defaultHeaders(addHIU, addHIP, addCM bool) map[string]string {
	headers := map[string]string{
		"Authorization": "Bearer " + c.AccessToken,
	}

	if addHIU {
		headers["X-HIU-ID"] = c.HIURegistrationID
	}
	if addHIP {
		headers["X-HIP-ID"] = c.HIPRegistrationID
	}
	if addCM {
		headers["X-CM-ID"] = c.ConsentManagerID
	}

	return headers
}
