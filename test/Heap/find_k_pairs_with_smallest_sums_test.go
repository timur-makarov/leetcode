package Heap

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Heap/find_k_pairs_with_smallest_sums"
	"testing"
)

func TestFindKPairswithSmallestSums(t *testing.T) {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 12}
	k := 3
	answer := [][]int{{1, 2}, {1, 4}, {7, 2}}
	result := find_k_pairs_with_smallest_sums.RunSolution(nums1, nums2, k)
	assert.ElementsMatch(t, answer, result)

	nums1 = []int{-10, -4, 0, 0, 6}
	nums2 = []int{3, 5, 6, 7, 8, 100}
	k = 10
	answer = [][]int{{-10, 3}, {-10, 5}, {-10, 6}, {-10, 7}, {-10, 8}, {-4, 3}, {-4, 5}, {-4, 6}, {-4, 7}, {0, 3}}
	result = find_k_pairs_with_smallest_sums.RunSolution(nums1, nums2, k)
	assert.ElementsMatch(t, answer, result)
}
