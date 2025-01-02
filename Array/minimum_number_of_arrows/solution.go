package minimum_number_of_arrows

import (
	"sort"
)

func RunSolution(points [][]int) int {
	sort.Slice(
		points, func(i, j int) bool {
			return points[i][0] < points[j][0]
		},
	)

	result := 1

	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			result++
		} else {
			points[i][0] = max(points[i-1][0], points[i][0])
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}

	return result
}
