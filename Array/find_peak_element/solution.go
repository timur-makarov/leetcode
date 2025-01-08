package find_peak_element

func RunSolution(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (start + end) / 2

		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}
