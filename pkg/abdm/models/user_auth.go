package models

import "time"

type UserAuthFetchModesRequest struct {
	RequestID string    `json:"requestId"`
	Timestamp time.Time `json:"timestamp"`
	Query     struct {
		UserAddress string          `json:"id"`
		Purpose     LinkPurposeType `json:"purpose"` // one of {"LINK", "KYC", "KYC_AND_LINK"}
		Requester   struct {
			Type  ServiceType `json:"type"`
			HIPID string      `json:"id"` // id of the HIP registered with ABDM
		} `json:"requester"`
	} `json:"query"`
}

type UserAuthFetchModesCallbackResponse struct {
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Auth      struct {
		Purpose string   `json:"purpose"`
		Modes   []string `json:"modes"`
	} `json:"auth"`
	Resp struct {
		RequestID string `json:"requestId"`
	} `json:"resp"`

	ErrorResponse `json:"error,omitempty"`
}

type UserAuthInitRequest struct {
	RequestID string    `json:"requestId"`
	Timestamp time.Time `json:"timestamp"`
	Query     struct {
		UserAddress string          `json:"id"`       // like devshoe@sbx
		Purpose     LinkPurposeType `json:"purpose"`  // one of {"LINK", "KYC", "KYC_AND_LINK"}
		AuthMode    AuthMode        `json:"authMode"` // one of {"MOBILE_OTP"}
		Requester   struct {
			Type  ServiceType `json:"type"`
			HIPID string      `json:"id"` // id of the HIP registered with ABDM
		} `json:"requester"`
	} `json:"query"`
}

type UserAuthInitCallbackResponse struct {
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Auth      struct {
		TransactionID string `json:"transactionId"`
		Mode          string `json:"mode"`
		Meta          struct {
			Hint   interface{} `json:"hint"`
			Expiry string      `json:"expiry"`
		} `json:"meta"`
	} `json:"auth"`
	Error ErrorResponse `json:"error,omitempty"`
	Resp  struct {
		RequestID string `json:"requestId"`
	} `json:"resp"`
}

type UserAuthConfirmRequest struct {
	RequestID     string `json:"requestId"`
	Timestamp     string `json:"timestamp"`
	TransactionID string `json:"transactionId"`
	Credential    struct {
		AuthCode string `json:"authCode"`
	} `json:"credential"`
}

type UserAuthConfirmCallbackResponse struct {
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Auth      struct {
		AccessToken string `json:"accessToken"`
		Patient     struct {
			ID           string `json:"id"` // address like devshoe@sbx
			Name         string `json:"name"`
			Gender       string `json:"gender"`
			YearOfBirth  int    `json:"yearOfBirth"`
			MonthOfBirth int    `json:"monthOfBirth"`
			DayOfBirth   int    `json:"dayOfBirth"`
			Address      struct {
				Line     interface{} `json:"line"` //todo: handle null without interface{}
				District string      `json:"district"`
				State    string      `json:"state"`
				Pincode  interface{} `json:"pincode"`
			} `json:"address"`
			Identifiers []struct {
				Type  string `json:"type"` // one of {"MOBILE", "HEALTH_NUMBER"}
				Value string `json:"value"`
			} `json:"identifiers"`
		} `json:"patient"`
	} `json:"auth"`
	Error ErrorResponse `json:"error,omitempty"`
	Resp  struct {
		RequestID string `json:"requestId"`
	} `json:"resp"`
}
