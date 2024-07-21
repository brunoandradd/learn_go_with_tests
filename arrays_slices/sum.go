package arraysslices

import "fmt"

func Sum(list []int) int {
	var total int
	for _, value := range list {
		total += value
	}
	return total
}

func SumAll(numbersTosum ...[]int) []int {
	var result []int

	for _, numbers := range numbersTosum {
		result = append(result, Sum(numbers))
	}
	return result
}

var puts = fmt.Println

func SumAllTails(numbersTosum ...[]int) []int {
	var result []int

	for _, numbers := range numbersTosum {
		if size := len(numbers); size == 0 {
			result = append(result, 0)
		} else {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}
	return result
}

func SumAllWithIndex(numbersTosum ...[]int) []int {
	arrayCount := len(numbersTosum)
	result := make([]int, arrayCount)

	for i := range arrayCount {
		result[i] = Sum(numbersTosum[i])
	}
	return result
}
