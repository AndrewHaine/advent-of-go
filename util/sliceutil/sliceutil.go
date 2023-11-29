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
