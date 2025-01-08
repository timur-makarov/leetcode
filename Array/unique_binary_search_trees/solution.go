package unique_binary_search_trees

func RunSolution(n int) int {
	if n < 2 {
		return 1
	}

	counts := make(map[int]int)

	counts[0] = 1
	counts[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			counts[i] += counts[j-1] * counts[i-j]
		}
	}

	return counts[n]
}
