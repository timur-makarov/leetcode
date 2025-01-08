package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/single_number"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	answer := 4
	result := single_number.RunSolution(nums)
	assert.Equal(t, answer, result)
}
