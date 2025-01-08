package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Heap/top_k_frequent_elements"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {
	nums := []int{4, 5, 5, 8, 8, 12, 15}
	k := 2
	answer := []int{5, 8}
	result := top_k_frequent_elements.RunSolution(nums, k)
	assert.ElementsMatch(t, answer, result)

	nums = []int{4, 5, 6, 7, 0, 0, 2}
	k = 1
	answer = []int{0}
	result = top_k_frequent_elements.RunSolution(nums, k)
	assert.ElementsMatch(t, answer, result)
}
