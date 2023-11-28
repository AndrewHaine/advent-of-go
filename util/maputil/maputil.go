package maputil

func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	inverted := map[V]K{}
	for k, v := range m {
		inverted[v] = k
	}
	return inverted
}
