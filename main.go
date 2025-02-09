package main

import (
	"github.com/timur-makarov/leetcode/Implementations/sorts"
)

func main() {
	sorts.RunSortFunctions()
}

func climbStairs(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2

	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
