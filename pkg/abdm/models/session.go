package models

// SessionRequest is the request body for authentication
// Requires ClientID and ClientSecret issued on the portal
type SessionRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// SessionResponse is the response body for authentication
// Returns AccessToken, RefreshToken, and ExpiresIn
// If error occurs, ErrorResponse is returned
type SessionResponse struct {
	AccessToken      string `json:"accessToken"`
	ExpiresIn        int    `json:"expiresIn"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresIn int    `json:"refreshExpiresIn"`
	TokenType        string `json:"tokenType"`

	ErrorResponse `json:"error,omitempty"`
}
