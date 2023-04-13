package models

import (
	"fmt"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) ToError() error {
	if e.Code == "" && e.Message == "" {
		return nil
	}
	return fmt.Errorf("%s: %s", e.Code, e.Message)
}
