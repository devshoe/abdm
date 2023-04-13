package abdm

var (
	URIBase    = "https://dev.abdm.gov.in/gateway"
	URISession = URIBase + "/v0.5/sessions" // POST
	URIBridge  = URIBase + "/bridges"       // PATCH

	URIRegisterService       = "https://dev.ndhm.gov.in/devservice/v1/bridges/services"    // PUT
	URIGetRegisteredServices = "https://dev.ndhm.gov.in/devservice/v1/bridges/getServices" // GET
)
