package models

type RegistrationType string

const (
	RegistrationTypeHIU RegistrationType = "HIU"
	RegistrationTypeHIP RegistrationType = "HIP"
)

type RegistrationRequest struct {
	ID     string           `json:"id"`
	Name   string           `json:"name"`
	Type   RegistrationType `json:"type"` // HIU or HIP
	Active bool             `json:"active"`
	Alias  []string         `json:"alias"`
}
