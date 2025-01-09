package maximum_number_of_distinct_elements_after_operations

import (
	"sort"
)

func RunSolution(nums []int, k int) int {
	sort.Ints(nums)

	result := 1
	last := nums[0] - k
	minPossibleVal := 0
	maxPossibleVal := 0

	for i := 1; i < len(nums); i++ {
		minPossibleVal = nums[i] - k
		maxPossibleVal = nums[i] + k

		if last < minPossibleVal {
			last = minPossibleVal
			result++
		} else if last < maxPossibleVal {
			last++
			result++
		}
	}

	return result
}
