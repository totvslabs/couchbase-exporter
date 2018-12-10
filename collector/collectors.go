package collector

const namespace = "couchbase"

func fromBool(i bool) float64 {
	if i {
		return 1.0
	}
	return 0.0
}

func last(ss []float64) float64 {
	if len(ss) == 0 {
		return 0.0
	}
	return ss[len(ss)-1]
}

// min returns the smaller of two given floats
func min(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}
