package harvest_api_client

func HaveSameFloat64Value(a, b *float64) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
