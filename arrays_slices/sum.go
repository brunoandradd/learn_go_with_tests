package arraysslices

func Sum(list []int) int {
	var total int
	for _, value := range list {
		total += value
	}
	return total
}
