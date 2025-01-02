package non_overlapping_intervals

import (
	"slices"
)

func RunSolution(intervals [][]int) int {
	slices.SortStableFunc(intervals, func(a, b []int) int { return a[1] - b[1] })

	result := 0
	prev := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[prev][1] <= intervals[i][0] {
			prev = i
		} else {
			result++
		}
	}

	return result
}
