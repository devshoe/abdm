package models

// RegistrationType is for enum on [HIU, HIP]
type RegistrationType string

// RegistrationType are the types of registration, either HIU or HIP
const (
	RegistrationTypeHIU RegistrationType = "HIU"
	RegistrationTypeHIP RegistrationType = "HIP"
)

// RegistrationRequest is the request body for registration
// It registers a service with ABDM, the ID must be unique
type RegistrationRequest struct {
	ID     string           `json:"id"`
	Name   string           `json:"name"`
	Type   RegistrationType `json:"type"` // HIU or HIP
	Active bool             `json:"active"`
	Alias  []string         `json:"alias"`
}

type RegistrationResponse struct {
	ErrorResponse `json:"error,omitempty"`
}

// RegisteredServicesResponse is the response body for get call
type RegisteredServicesResponse struct {
	Bridge struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		URL         string `json:"url"`
		Active      bool   `json:"active"`
		Blocklisted bool   `json:"blocklisted"`
	} `json:"bridge"`
	Services []struct {
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		Types     []string `json:"types"`
		Endpoints struct {
		} `json:"endpoints"`
		Active bool `json:"active"`
	} `json:"services"`

	ErrorResponse `json:"error,omitempty"`
}
