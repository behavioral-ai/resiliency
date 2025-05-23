package test

import (
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
	"github.com/behavioral-ai/intermediary/cache"
	"github.com/behavioral-ai/intermediary/cache/cachetest"
	"github.com/behavioral-ai/intermediary/config"
	urn2 "github.com/behavioral-ai/intermediary/urn"
	"github.com/behavioral-ai/resiliency/operations"
	"github.com/behavioral-ai/traffic/limiter"
	"github.com/behavioral-ai/traffic/redirect"
	"github.com/behavioral-ai/traffic/urn"
)

func NewRootEndpoint() *rest.Endpoint {
	_ = operations.Agent
	_ = cache.NamespaceName
	_ = limiter.NamespaceName
	_ = redirect.NamespaceName
	cache := host.Agent(urn2.CacheAgent)
	cache.Message(httpx.NewConfigExchangeMessage(cachetest.Exchange))
	m := make(map[string]string)
	m[config.CacheHostKey] = "localhost:8082"
	cache.Message(messaging.NewConfigMapMessage(m))
	/*
		chain := httpx.BuildChain(host.AccessLogLink, host.AuthorizationLink,
			exchange.Agent(urn.RedirectAgent),
			exchange.Agent(urn2.CacheAgent),
			exchange.Agent(urn.LimiterAgent), RoutingLink)

	*/

	return host.NewEndpoint(host.Agent(urn.RedirectAgent),
		host.Agent(urn2.CacheAgent),
		host.Agent(urn.LimiterAgent), RoutingLink)
}
