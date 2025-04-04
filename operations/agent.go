package operations

import (
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/intermediary/cache"
	"github.com/behavioral-ai/intermediary/routing"
	"github.com/behavioral-ai/traffic/analytics"
	"github.com/behavioral-ai/traffic/limiter"
	"github.com/behavioral-ai/traffic/redirect"
)

const (
	NamespaceName = "resiliency:agent/behavioral-ai/resiliency/operations"
)

// TODO : need host name
type agentT struct {
	running bool

	agents     *messaging.Exchange
	notifier   eventing.NotifyFunc
	dispatcher eventing.Dispatcher
}

// New - create a new operations agent
func New(notifier eventing.NotifyFunc) messaging.Agent {
	return newAgent(notifier, nil)
}

func newAgent(notifier eventing.NotifyFunc, dispatcher eventing.Dispatcher) *agentT {
	a := new(agentT)

	a.agents = messaging.NewExchange()
	a.agents.RegisterMailbox(cache.Agent)
	a.agents.RegisterMailbox(limiter.Agent)
	a.agents.RegisterMailbox(analytics.Agent)
	a.agents.RegisterMailbox(redirect.Agent)
	a.agents.RegisterMailbox(routing.Agent)

	if notifier == nil {
		a.notifier = eventing.OutputNotify
	} else {
		a.notifier = notifier
	}
	if dispatcher == nil {
		a.dispatcher = eventing.NewTraceDispatcher()
	} else {
		a.dispatcher = dispatcher
	}
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Event() == messaging.ConfigEvent {
		a.configure(m)
		return
	}
	switch m.Event() {
	case messaging.StartupEvent:
		a.agents.Broadcast(m)
	case messaging.ShutdownEvent:
		a.agents.Broadcast(m)
	case messaging.PauseEvent:
		a.agents.Broadcast(m)
	case messaging.ResumeEvent:
		a.agents.Broadcast(m)
	case eventing.NotifyEvent:
		a.notifier(eventing.NotifyContent(m))
	case eventing.DispatchEvent:
		i := eventing.DispatchContent(m)
		a.dispatcher.Dispatch(a, i.Channel, i.Event)
	case eventing.ActivityEvent:
		/*
			if m.ContentType() == ContentTypeActivity {
				a.addActivity(ActivityContent(m))
				return
			}
		*/
	}
}

func (a *agentT) dispatch(channel any, event1 string) {
	//a.handler.Message(event.NewDispatchMessage(a, channel, event1))
}

func (a *agentT) shutdown() {
	a.agents.Shutdown()
}

func (a *agentT) configure(m *messaging.Message) {
	if n, ok := configNotifierContent(m); ok {
		a.notifier = n
	}
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}
