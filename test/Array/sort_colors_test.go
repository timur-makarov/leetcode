package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/sort_colors"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	answer := []int{0, 0, 1, 1, 2, 2}
	sort_colors.RunSolution(nums)
	assert.Equal(t, answer, nums)
}
