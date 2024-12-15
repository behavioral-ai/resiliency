package knowledge

type Frame struct {
	Observe *Observation
	Impress Impression
	Act     Action
}

func NewFrame(o *Observation) *Frame {
	f := new(Frame)
	f.Observe = o
	return f
}
