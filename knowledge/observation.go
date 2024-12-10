package knowledge

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"reflect"
)

const (
	ContentTypeObservation = "application/observation"
)

type Observation struct {
	Latency  int
	Gradient int
}

func GetObservation(h messaging.Notifier, agentId string, msg *messaging.Message) (Observation, *core.Status) {
	if !msg.IsContentType(ContentTypeObservation) {
		return Observation{}, core.StatusNotFound()
	}
	if p, ok := msg.Body.(Observation); ok {
		return p, core.StatusOK()
	}
	status := observationTypeErrorStatus(agentId, msg.Body)
	h.Notify(status)
	return Observation{}, status
}

func observationTypeErrorStatus(agentId string, t any) *core.Status {
	err := errors.New(fmt.Sprintf("error: observation type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}
