package abdm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func do(c *http.Client, method string, url string, headers map[string]string, requestBody any, responseReader any) (code int, err error) {
	var (
		bodyb    []byte
		request  *http.Request
		response *http.Response
	)

	if bodyb, err = marshal(requestBody); err != nil {
		return -1, fmt.Errorf("do: %w", err)
	} else if request, err = createRequest(method, url, bodyb, headers); err != nil {
		return -1, fmt.Errorf("do: %w", err)
	} else if response, err = c.Do(request); err != nil {
		return response.StatusCode, fmt.Errorf("do: send request - %w", err)
	} else if err = unmarshal(response.Body, responseReader); err != nil {
		return response.StatusCode, fmt.Errorf("do: %w", err)
	}

	return response.StatusCode, nil
}

func marshal(item any) (itemb []byte, err error) {
	switch item := item.(type) {
	case string:
		itemb = []byte(item)
	case []byte:
		itemb = item
	default:
		if itemb, err = json.Marshal(item); err != nil {
			return nil, fmt.Errorf("marshal body - %w", err)
		}
	}
	return
}

func createRequest(method string, url string, body []byte, headers ...map[string]string) (request *http.Request, err error) {
	if body == nil {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return nil, fmt.Errorf("request create - %w", err)
	}

	if len(headers) > 0 && headers[0] != nil {
		for k, v := range headers[0] {
			request.Header.Set(k, v)
		}
	}

	return
}

func unmarshal(item io.Reader, itemReader any) (err error) {
	if itemReader != nil {
		if err = json.NewDecoder(item).Decode(itemReader); err != nil {
			return fmt.Errorf("unmarshal body - %w", err)
		}
	}
	return
}
