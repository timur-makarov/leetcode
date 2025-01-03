package three_sum

import "slices"

func RunSolution(nums []int) [][]int {
	slices.Sort(nums)

	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		targetSum := -nums[i]
		wStart := i + 1
		wEnd := len(nums) - 1

		for wStart < wEnd {
			sum := nums[wStart] + nums[wEnd]

			if sum == targetSum {
				result = append(result, []int{-targetSum, nums[wStart], nums[wEnd]})
				wStart++
				wEnd--

				for wStart < wEnd && nums[wStart] == nums[wStart-1] {
					wStart++
				}

				for wStart < wEnd && nums[wEnd] == nums[wEnd+1] {
					wEnd--
				}
			} else if targetSum > sum {
				wStart++
			} else {
				wEnd--
			}
		}
	}

	return result
}
