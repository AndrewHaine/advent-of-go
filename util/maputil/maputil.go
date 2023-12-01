package maputil

import (
	"sort"
)

func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	inverted := map[V]K{}
	for k, v := range m {
		inverted[v] = k
	}
	return inverted
}

func Values[K comparable, V comparable](m map[K]V) []V {
	values := []V{}
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func SortedValues[V comparable](m map[int]V) []V {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	values := make([]V, 0, len(m))
	for _, k := range keys {
		values = append(values, m[k])
	}
	return values
}
