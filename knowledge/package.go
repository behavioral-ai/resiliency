package knowledge

import (
	"context"
	"fmt"
	"github.com/behavioral-ai/core/core"
)

const (
	PkgPath    = "github/behavioral-ai/resiliency/knowledge"
	WestRegion = "us-west1"
	WestZoneA  = "w-a"
	WestZoneB  = "w-b"

	CentralRegion = "us-central1"
	CentralZoneA  = "c-a"
	CentralZoneB  = "c-b"

	EastRegion = "us-east1"
	EastZoneA  = "e-a"
	EastZoneB  = "e-b"
)

func AddAction(ctx context.Context, origin core.Origin, action int) *core.Status {
	fmt.Printf("Action: %v %v\n", origin, action)
	return core.StatusOK()
}

func AddInference(ctx context.Context, origin core.Origin, f *Frame) *core.Status {
	if f != nil {
		fmt.Printf("Frame: %v %v\n", origin, f.Observe)
	}
	return core.StatusOK()
}

func Inference(o *Observation) *Frame {
	f := NewFrame(o)
	//if f == nil {
	//	return
	//	}
	comprehension(f)
	reasoning(f)
	return f
}
