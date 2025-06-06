package utils

func FilterArray[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func ExistsInArray[T comparable](slice []T, seek T) bool {
	for _, val := range slice {
		if val == seek {
			return true
		}
	}
	return false
}
