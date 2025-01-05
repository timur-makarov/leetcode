package subsets_ii

import (
	"slices"
)

func RunSolution(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})

	slices.Sort(nums)

	prevNum := 0
	toIndex := 0

	for i, num := range nums {
		var changes [][]int
		fromIndex := 0

		if i > 0 && prevNum == num {
			fromIndex = toIndex + 1
		}

		toIndex = len(result) - 1

		for j := fromIndex; j <= toIndex; j++ {
			changes = append(changes, append(slices.Clone(result[j]), num))
		}

		result = append(result, changes...)
		prevNum = num
	}

	return result
}
