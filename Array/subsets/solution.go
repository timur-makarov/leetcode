package subsets

import (
	"math"
	"slices"
)

func RunSolution(nums []int) [][]int {
	resultCap := math.Pow(2.0, float64(len(nums)))
	result := make([][]int, 0, int(resultCap))
	result = append(result, []int{})

	for _, num := range nums {
		var changes [][]int
		for _, item := range result {
			changes = append(changes, append(slices.Clone(item), num))
		}
		result = append(result, changes...)
	}

	return result
}
