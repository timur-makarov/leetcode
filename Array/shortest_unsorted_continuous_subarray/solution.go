package shortest_unsorted_continuous_subarray

import (
	"slices"
)

func RunSolution(nums []int) int {
	wStart := 0
	wEnd := len(nums) - 1

	for wStart < wEnd && nums[wStart] <= nums[wStart+1] {
		wStart++
	}

	if wStart == wEnd {
		return 0
	}

	for wEnd > wStart && nums[wEnd] >= nums[wEnd-1] {
		wEnd--
	}

	minNum := slices.Min(nums[wStart : wEnd+1])
	maxNum := slices.Max(nums[wStart : wEnd+1])

	for wStart > 0 && nums[wStart-1] > minNum {
		wStart--
	}

	for wEnd < len(nums)-1 && nums[wEnd+1] < maxNum {
		wEnd++
	}

	return wEnd - wStart + 1
}
