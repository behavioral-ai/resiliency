package knowledge

// observation -> comprehend -> impression
// impression -> reason -> action

const (
	threshold = 2500
)

type lookup struct {
	low    int
	medium int
	high   int
}

var (
	sat  = lookup{low: 25, medium: 65, high: 80}
	grad = lookup{low: 5, medium: 10, high: 15}
)

func comprehension(f *Frame) {
	if f == nil {
		return
	}
	f.Impress.Saturation = lookupSaturation(f.Observe.Latency)
	f.Impress.Gradient = lookupGradient(f.Observe.Gradient)
}

func lookupSaturation(latency int) string {
	s := float64(latency) / float64(threshold)
	s *= 100
	if s <= float64(sat.low) {
		return low
	}
	if s <= float64(sat.medium) {
		return medium
	}
	return high
}

func lookupGradient(gradient int) string {
	if gradient <= grad.low {
		return low
	}
	if gradient <= grad.medium {
		return medium
	}
	return high
}

func reasoning(f *Frame) {
	if f == nil {
		return
	}
}
