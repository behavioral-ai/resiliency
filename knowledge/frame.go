package knowledge

const (
	low    = "low"
	medium = "medium"
	high   = "high"
)

type Impression struct {
	Saturation string
	Gradient   string
}

type Frame struct {
	Observe Observation
	Impress Impression
	Action  int
	Stars   int
}

func NewFrame(o Observation) *Frame {
	f := new(Frame)
	f.Observe = o
	return f
}
