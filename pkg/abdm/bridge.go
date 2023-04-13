package abdm

import (
	"fmt"

	"github.com/devshoe/abdm/pkg/abdm/models"
)

// UpdateBridgeURL updates the bridge URL registered with ABDM. i.e where our service can be reached
func (c *Client) UpdateBridgeURL(url string) (err error) {
	var (
		request  models.BridgeRequest
		response models.BridgeResponse
		method                     = "PATCH"
		headers  map[string]string = c.defaultHeaders(false, false, false)
		code     int
	)

	request.CallbackURL = url

	if code, err = do(c.c, method, URISession, headers, request, &response); err != nil {
		return err
	} else if code != 200 {
		return fmt.Errorf("register bridge url: response code %d", code)
	}

	return
}
