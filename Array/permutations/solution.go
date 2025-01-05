package permutations

import "slices"

func RunSolution(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	var result [][]int

	for i := 0; i < len(nums); i++ {
		v := slices.Clone(nums)

		v[i], v[0] = v[0], v[i]

		v = v[1:]

		res := RunSolution(v)

		for _, j := range res {
			j = append([]int{nums[i]}, j...)
			result = append(result, j)
		}
	}

	return result
}
