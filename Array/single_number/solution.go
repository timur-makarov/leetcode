package single_number

func RunSolution(nums []int) int {
	result := 0

	for _, num := range nums {
		result ^= num
	}

	return result
}
