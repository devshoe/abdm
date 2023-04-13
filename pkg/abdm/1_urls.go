package abdm

var (
	URIBase    = "https://dev.abdm.gov.in/gateway"
	URISession = URIBase + "/v0.5/sessions" // POST
	URIBridge  = URIBase + "/bridges"       // PATCH

	URIRegisterService       = "https://dev.ndhm.gov.in/devservice/v1/bridges/services"      // PUT
	URIGetRegisteredServices = "https://dev.ndhm.gov.in/devservice/v1/bridges/getServices"   // GET
	URIUserAuthFetchModes    = "https://dev.ndhm.gov.in/gateway/v0.5/users/auth/fetch-modes" // POST
	URIUserAuthInit          = "https://dev.ndhm.gov.in/gateway/v0.5/users/auth/init"        // POST
	URIUserAuthConfirm       = "https://dev.ndhm.gov.in/gateway/v0.5/users/auth/confirm"     // POST
)
