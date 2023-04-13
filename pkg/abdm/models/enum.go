package models

// ServiceType is for enum on allowed service types in registration requests
type ServiceType string

// ServiceType are the types of registration, either HIU or HIP
const (
	ServiceTypeHIU ServiceType = "HIU"
	ServiceTypeHIP ServiceType = "HIP"
)

// LinkPurposeType is for enum on allowed purpose types in user auth requests
type LinkPurposeType string

// LinksPurposeType are the types of purpose, either LINK, KYC or KYC_AND_LINK
const (
	LinkPurposeTypeLink       LinkPurposeType = "LINK"
	LinkPurposeTypeKYC        LinkPurposeType = "KYC"
	LinkPurposeTypeLinkAndKYC LinkPurposeType = "KYC_AND_LINK"
)

// AuthMode is for enum on allowed auth modes in user auth requests
type AuthMode string

// AuthModes are the types of auth modes, either MOBILE_OTP
const (
	AuthModeMobileOTP AuthMode = "MOBILE_OTP"
)
