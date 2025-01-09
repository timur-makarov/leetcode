package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/find_k_closest_elements"
	"testing"
)

func TestFindKClosestElements(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 12, 15}
	k := 3
	x := 13
	answer := []int{8, 12, 15}
	result := find_k_closest_elements.RunSolution(nums, k, x)
	assert.Equal(t, answer, result)

	nums = []int{0, 1, 2, 4, 5, 6, 7}
	k = 2
	x = 0
	answer = []int{0, 1}
	result = find_k_closest_elements.RunSolution(nums, k, x)
	assert.Equal(t, answer, result)
}
