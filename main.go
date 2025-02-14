package main

import (
	"log"
	"slices"
)

func main() {
	//sorts.RunSortFunctions()

	nums := []int{0, 1, 0, 3, 2, 3}
	log.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	memo := make([]int, len(nums))

	for i := range memo {
		memo[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] < nums[j] {
				memo[i] = max(memo[i], memo[j]+1)
			}
		}
	}

	log.Println(memo)

	return slices.Max(memo)
}
