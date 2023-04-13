package models

// AuthenticationRequest is the request body for authentication
// Requires ClientID and ClientSecret issued on the portal
type AuthenticationRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// AuthenticationResponse is the response body for authentication
// Returns AccessToken, RefreshToken, and ExpiresIn
// If error occurs, ErrorResponse is returned
type AuthenticationResponse struct {
	AccessToken      string `json:"accessToken"`
	ExpiresIn        int    `json:"expiresIn"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresIn int    `json:"refreshExpiresIn"`
	TokenType        string `json:"tokenType"`

	ErrorResponse `json:"error,omitempty"`
}
