package sliceutil

func Chunk[T comparable](slice []T, size int) [][]T {
	chunks := [][]T{}
	chunk := []T{}
	for _, item := range slice {
		chunk = append(chunk, item)
		if len(chunk) == size {
			chunks = append(chunks, chunk)
			chunk = []T{}
		}
	}
	if len(chunk) != 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func Pop[T comparable](slice []T) (T, []T) {
	length := len(slice)
	elem := slice[length-1]
	slice = slice[:length-1]
	return elem, slice
}

func Every[T comparable](slice []T, callback func(index int, item T) bool) (result bool) {
	result = true
	for i, item := range slice {
		if !callback(i, item) {
			result = false
			break
		}
	}
	return
}

func Some[T comparable](slice []T, callback func(index int, item T) bool) (result bool) {
	result = false
	for i, item := range slice {
		if callback(i, item) {
			result = true
			break
		}
	}
	return
}
