package endpoint

import (
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	"github.com/behavioral-ai/intermediary/cache"
	"github.com/behavioral-ai/intermediary/cache/cachetest"
	"github.com/behavioral-ai/intermediary/routing"
	"github.com/behavioral-ai/intermediary/routing/routingtest"
	urn2 "github.com/behavioral-ai/intermediary/urn"
	link2 "github.com/behavioral-ai/resiliency/link"
	"github.com/behavioral-ai/traffic/limiter"
	"github.com/behavioral-ai/traffic/redirect"
	"github.com/behavioral-ai/traffic/urn"
)

func newRootEndpoint() *rest.Endpoint {
	_ = cache.NamespaceName
	_ = routing.NamespaceName
	_ = limiter.NamespaceName
	_ = redirect.NamespaceName

	// Overriding agent http exchange
	exchange.Agent(urn2.CacheAgent).Message(httpx.NewConfigExchangeMessage(cachetest.Exchange))
	exchange.Agent(urn2.RoutingAgent).Message(httpx.NewConfigExchangeMessage(routingtest.Exchange))

	return host.NewEndpoint(link2.Logger,
		exchange.Agent(urn.RedirectAgent),
		exchange.Agent(urn2.CacheAgent),
		exchange.Agent(urn.LimiterAgent),
		exchange.Agent(urn2.RoutingAgent))
}

//cacheAgent := exchange.Agent(urn2.CacheAgent)
//cacheAgent.Message(httpx.NewConfigExchangeMessage(cachetest.Exchange))
//m := make(map[string]string)
//m[config.CacheHostKey] = "localhost:8082"
//cacheAgent.Message(messaging.NewConfigMapMessage(m))

//routingAgent := exchange.Agent(urn2.RoutingAgent)
//routingAgent.Message(httpx.NewConfigExchangeMessage(routingtest.Exchange))
//m[config.AppHostKey] = "localhost:8080"
//m[config.TimeoutKey] = "10ms"
//routingAgent.Message(messaging.NewConfigMapMessage(m))
