package models

// RegistrationRequest is the request body for registration
// It registers a service with ABDM, the ID must be unique
type RegistrationRequest struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Type   ServiceType `json:"type"` // HIU or HIP
	Active bool        `json:"active"`
	Alias  []string    `json:"alias"`
}

type RegistrationResponse struct {
	ErrorResponse `json:"error,omitempty"`
}

// RegisteredServicesResponse is the response body for get call
type RegisteredServicesResponse struct {
	Bridge   BridgeInfo    `json:"bridge"`
	Services []ServiceInfo `json:"services"`

	ErrorResponse `json:"error,omitempty"`
}

// BridgeInfo is data pertaining to the bridge, like id registered on portal and callback url
type BridgeInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Active      bool   `json:"active"`
	Blocklisted bool   `json:"blocklisted"`
}

// ServiceInfo is data pertaining to a service, like id used for registering and type
type ServiceInfo struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Types     []string `json:"types"`
	Endpoints struct {
	} `json:"endpoints"`
	Active bool `json:"active"`
}
