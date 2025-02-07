package partition_equal_subset_sum

func RunSolution(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}
	hs := sum / 2

	memo := make([]bool, hs+1)
	memo[0] = true

	for _, num := range nums {
		for s := hs; s >= num; s-- {
			memo[s] = memo[s] || memo[s-num]
		}
	}

	return memo[hs]
}
