package collector

const globalNamespace = "couchbase"

func fromBool(i bool) float64 {
	if i {
		return 1.0
	}
	return 0.0
}
