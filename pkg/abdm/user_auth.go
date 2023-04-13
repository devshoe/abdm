package abdm

import (
	"fmt"
	"time"

	"github.com/devshoe/abdm/pkg/abdm/models"
	"github.com/google/uuid"
)

// LinkPurposeIsAllowedForUser returns 202 status if ok
// todo: doc says something else, it should be returning allowed auth modes :/
func (c *Client) LinkPurposeIsAllowedForUser(useraddress string, purpose models.LinkPurposeType) (err error) {
	var (
		request  models.UserAuthFetchModesRequest
		response models.UserAuthFetchModesResponse
		method                     = "POST"
		headers  map[string]string = c.defaultHeaders(false, true, false)
		code     int
	)

	request.RequestID = uuid.NewString()
	request.Timestamp = time.Now()
	request.Query.UserAddress = useraddress
	request.Query.Purpose = purpose
	request.Query.Requester.Type = models.ServiceTypeHIP
	request.Query.Requester.HIPID = c.HIPRegistrationID

	if code, err = do(c.c, method, URISession, headers, request, &response); err != nil {
		return err
	} else if code != 202 {
		return fmt.Errorf("user auth fetch modes: response code %d", code)
	}

	return
}
