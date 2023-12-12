package mathutil

func SumInts(in []int) (sum int) {
	sum = 0
	for _, num := range in {
		sum += num
	}
	return
}
