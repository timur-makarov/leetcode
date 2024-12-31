package merge_intrevals

import "slices"

func RunSolution(intervals [][]int) [][]int {
	var result [][]int

	if len(intervals) < 2 {
		return intervals
	}

	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })

	cStart := intervals[0][0]
	cEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] <= cEnd {
			cEnd = max(interval[1], cEnd)
		} else {
			result = append(result, []int{cStart, cEnd})
			cStart = interval[0]
			cEnd = interval[1]
		}
	}

	result = append(result, []int{cStart, cEnd})

	return result
}
