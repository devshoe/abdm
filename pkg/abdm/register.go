package abdm

import "github.com/devshoe/abdm/pkg/abdm/models"

// RegisterHIU registers a HIU with ABDM. The uid is the unique identifier for the HIU
// and it will be sent as header `X-HIU-ID` in any subsequent request
func (c *Client) RegisterHIU(uid, name string, aliases []string) (err error) {
	var (
		request  models.RegistrationRequest
		response models.RegistrationResponse
		method   = "PUT"
		headers  = c.defaultHeaders(false, false, false)
	)

	request.ID = uid
	request.Name = name
	request.Type = models.RegistrationTypeHIU
	request.Active = true
	request.Alias = aliases

	if _, err = do(c.c, method, URIRegisterService, headers, request, nil); err != nil {
		return err
	}

	return response.ToError()
}

// RegisterHIP registers a HIP with ABDM. The uid is the unique identifier for the HIP
// and it will be sent as header `X-HIP-ID` in any subsequent request
func (c *Client) RegisterHIP(uid, name string, aliases []string) (err error) {
	var (
		request  models.RegistrationRequest
		response models.RegistrationResponse
		method   = "PUT"
		headers  = c.defaultHeaders(false, false, false)
	)

	request.ID = uid
	request.Name = name
	request.Type = models.RegistrationTypeHIU
	request.Active = true
	request.Alias = aliases

	if _, err = do(c.c, method, URIRegisterService, headers, request, nil); err != nil {
		return err
	}

	return response.ToError()
}

// GetRegisteredServices returns a list of all registered services which were created using `RegisterHIP` or `RegisterHIU`
func (c *Client) GetRegisteredServices() (models.RegisteredServicesResponse, error) {
	var (
		response models.RegisteredServicesResponse
		method   = "GET"
		headers  = c.defaultHeaders(false, false, false)

		err error
	)

	if _, err = do(c.c, method, URIGetRegisteredServices, headers, nil, &response); err != nil {
		return response, err
	}

	return response, response.ToError()
}
