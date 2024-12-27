package minimum_size_subarray_sum

func RunSolution(target int, nums []int) int {
	size := len(nums)
	minLength := size + 1
	wSum := 0
	wStart := 0

	for wEnd := 0; wEnd < size; wEnd++ {
		wSum += nums[wEnd]

		for wSum >= target {
			minLength = min(minLength, wEnd-wStart+1)
			wSum -= nums[wStart]
			wStart++
		}
	}

	if minLength == size+1 {
		return 0
	}

	return minLength
}
