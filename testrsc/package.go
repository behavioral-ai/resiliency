package testrsc

const (
	PkgPath = "github/behavioral-ai/resiliency/testrsc"
)
const (
	commonBasePath = "file:///f:/files/common"

	NotFoundResp            = commonBasePath + "/http-404-resp.txt"
	OKResp                  = commonBasePath + "/http-200-resp.txt"
	GatewayTimeoutResp      = commonBasePath + "/http-504-resp.txt"
	InternalServerErrorResp = commonBasePath + "/internal-server-error-resp.txt"
)

const (
	log1BasePath = "file:///f:/files/event1"

	LOG1IngressEntry          = log1BasePath + "/ingress-entry.json"
	LOG1IngressGetAllReq      = log1BasePath + "/ingress-get-all-req.txt"
	LOG1IngressGetAllResp     = log1BasePath + "/ingress-get-all-resp.txt"
	LOG1IngressPutReq         = log1BasePath + "/ingress-put-req.txt"
	LOG1IngressGetNotFoundReq = log1BasePath + "/ingress-get-not-found-req.txt"
	log2BasePath              = "file:///f:/files/event2"

	LOG2IngressEntry      = log2BasePath + "/ingress-entry.json"
	LOG2IngressGetAllReq  = log2BasePath + "/ingress-get-all-req.txt"
	LOG2IngressGetAllResp = log2BasePath + "/ingress-get-all-resp.txt"

	//
)
