package Graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Graph/path_sum_iii"
)

func TestPathSumIII(t *testing.T) {
	var tr Tree

	_ = tr.Insert(9)
	_ = tr.Insert(3)
	_ = tr.Insert(20)
	_ = tr.Insert(15)
	_ = tr.Insert(10)
	_ = tr.Insert(25)
	_ = tr.Insert(1)
	_ = tr.Insert(4)

	targetSum := 25
	answer := 2
	result := path_sum_iii.RunSolution(tr.Root, targetSum)
	assert.Equal(t, answer, result)
}
