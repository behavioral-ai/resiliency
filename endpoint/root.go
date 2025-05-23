package endpoint

import (
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	_ "github.com/behavioral-ai/intermediary/cache"
	"github.com/behavioral-ai/intermediary/cache/cachetest"
	_ "github.com/behavioral-ai/intermediary/routing"
	"github.com/behavioral-ai/intermediary/routing/routingtest"
	urn2 "github.com/behavioral-ai/intermediary/urn"
	link "github.com/behavioral-ai/resiliency/link"
	_ "github.com/behavioral-ai/traffic/limiter"
	_ "github.com/behavioral-ai/traffic/redirect"
	"github.com/behavioral-ai/traffic/urn"
)

func newRootEndpoint() *rest.Endpoint {
	// Overriding agent http exchange
	host.Agent(urn2.CacheAgent).Message(httpx.NewConfigExchangeMessage(cachetest.Exchange))
	host.Agent(urn2.RoutingAgent).Message(httpx.NewConfigExchangeMessage(routingtest.Exchange))

	return host.NewEndpoint(link.Logger,
		host.Agent(urn.RedirectAgent),
		host.Agent(urn2.CacheAgent),
		host.Agent(urn.LimiterAgent),
		host.Agent(urn2.RoutingAgent))
}
