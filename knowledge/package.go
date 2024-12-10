package knowledge

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

func Comprehend(o *Observation) *Impression {
	i := new(Impression)

	return i
}

func Reason(i *Impression) *Action {
	a := new(Action)
	return a
}
