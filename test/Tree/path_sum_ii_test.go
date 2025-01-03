package Tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Tree/path_sum_ii"
	"testing"
)

func TestPathSumII(t *testing.T) {
	var tr Tree

	_ = tr.Insert(9)
	_ = tr.Insert(3)
	_ = tr.Insert(20)
	_ = tr.Insert(15)
	_ = tr.Insert(10)
	_ = tr.Insert(25)
	_ = tr.Insert(1)
	_ = tr.Insert(4)

	answer := [][]int{{9, 20, 15, 10}, {9, 20, 25}}
	result := path_sum_ii.RunSolution(tr.Root, 54)
	assert.Equal(t, answer, result)
}
