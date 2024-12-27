package longest_subarray_of_1s_after_deleting_one_element

func RunSolution(nums []int) int {
	wStart := 0
	maxLength := 0
	zeros := 0

	for wEnd := 0; wEnd < len(nums); wEnd++ {
		if nums[wEnd] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[wEnd] == 0 {
				zeros--
			}
			wStart++
		}

		maxLength = max(maxLength, wEnd-wStart-zeros)
	}

	if maxLength == len(nums) {
		return maxLength + 1
	}

	return maxLength
}
