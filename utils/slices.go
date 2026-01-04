package utils

func SliceContainsFunc[T any](slice []T, fn func(T) bool) bool {
	for _, item := range slice {
		if fn(item) {
			return true
		}
	}
	return false
}
