package models

type EndpointDescription[request any, response any] struct {
	URL          string            `json:"url"`
	Method       string            `json:"method"`
	Headers      map[string]string `json:"headers"`
	RequestBody  request           `json:"request_body"`
	ResponseBody response          `json:"response_body"`
	ExpectStatus int               `json:"expect_status"`
}
