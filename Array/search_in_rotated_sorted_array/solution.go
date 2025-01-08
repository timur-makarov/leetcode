package search_in_rotated_sorted_array

func RunSolution(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	pivot := start
	start = 0
	end = len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		realMid := (mid + pivot) % len(nums)

		if nums[realMid] == target {
			return realMid
		}

		if nums[realMid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return -1
}
