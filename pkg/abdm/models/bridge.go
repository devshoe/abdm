package models

// BridgeRequest is the PATCH request body for bridge
// It updates the bridge URL in the gateway, ie where our service can be reached
type BridgeRequest struct {
	CallbackURL string `json:"url"`
}

// BridgeResponse is the response body for bridge
// Only contains data if error occurs
type BridgeResponse struct {
	ErrorResponse `json:"error,omitempty"`
}
