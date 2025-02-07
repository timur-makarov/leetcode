package knapsack

func RunSolution(profits, weights []int, cap int) int {
	var memo [][]int

	for range 2 {
		memo = append(memo, make([]int, cap+1))
	}

	for c := weights[0]; c <= cap; c++ {
		memo[0][c] = profits[0]
	}

	for i := 1; i < len(profits); i++ {
		for c := 1; c <= cap; c++ {
			if weights[i] <= c {
				memo[i%2][c] = profits[i] + memo[(i-1)%2][c-weights[i]]
			} else {
				memo[i%2][c] = memo[(i-1)%2][c]
			}
		}
	}

	return memo[(len(profits)-1)%2][cap]
}
