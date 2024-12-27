package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/fruits_into_baskets"
	"testing"
)

func TestFruitsintoBaskets(t *testing.T) {
	fruits := []int{1, 2, 3, 2, 2}
	answer := 4
	result := fruits_into_baskets.RunSolution(fruits)
	assert.Equal(t, answer, result)
}
