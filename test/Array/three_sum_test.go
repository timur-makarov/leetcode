package Array

import (
	"github.com/stretchr/testify/assert"
	three_sum "github.com/timur-makarov/leetcode/Array/3sum"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	answer := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result := three_sum.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{0, 1, 1}
	answer = [][]int{}
	result = three_sum.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{0, 0, 0}
	answer = [][]int{{0, 0, 0}}
	result = three_sum.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{1, 2, -2, -1}
	answer = [][]int{}
	result = three_sum.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{-2, 0, 0, 2, 2}
	answer = [][]int{{-2, 0, 2}}
	result = three_sum.RunSolution(nums)
	assert.Equal(t, answer, result)
}
