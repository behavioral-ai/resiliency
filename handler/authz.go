package operations

import (
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/core/rest"
	"net/http"
)

const (
	NamespaceNameAuth = "test:resiliency:handler/authorization/http"
	AuthorizationName = "Authorization"
)

func init() {
	exchange.RegisterExchangeHandler(NamespaceNameAuth, Authorization)
}

func Authorization(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(AuthorizationName)
		if auth == "" {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
		return next(r)
	}
}
