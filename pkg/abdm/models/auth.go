package models

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
