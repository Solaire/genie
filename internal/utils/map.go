package utils

func MapGetDefault[K comparable, V any](m map[K]V, key K, _default V) V {
	if val, ok := m[key]; ok {
		return val
	}
	return _default
}

func ExistsInMap[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}
