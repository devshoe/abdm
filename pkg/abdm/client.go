package abdm

import (
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	c            *http.Client
	ClientID     string
	ClientSecret string
	AccessToken  string
	RefreshToken string
}

func New(clientID, clientSecret string) *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		c:            &http.Client{Jar: jar},
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}
