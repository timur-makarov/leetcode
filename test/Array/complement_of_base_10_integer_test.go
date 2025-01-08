package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/complement_of_base_10_integer"
	"testing"
)

func TestComplementofBase10Integer(t *testing.T) {
	num := 10
	answer := 5
	result := complement_of_base_10_integer.RunSolution(num)
	assert.Equal(t, answer, result)

	num = 7
	answer = 0
	result = complement_of_base_10_integer.RunSolution(num)
	assert.Equal(t, answer, result)

	num = 5
	answer = 2
	result = complement_of_base_10_integer.RunSolution(num)
	assert.Equal(t, answer, result)
}
