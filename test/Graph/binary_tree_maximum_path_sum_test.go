package Graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Graph/binary_tree_maximum_path_sum"
)

func TestBinaryTreeMaximumPathSum(t *testing.T) {
	var tr Tree

	_ = tr.Insert(9)
	_ = tr.Insert(3)
	_ = tr.Insert(20)
	_ = tr.Insert(15)
	_ = tr.Insert(10)
	_ = tr.Insert(25)
	_ = tr.Insert(1)
	_ = tr.Insert(4)

	answer := 70
	result := binary_tree_maximum_path_sum.RunSolution(tr.Root)
	assert.Equal(t, answer, result)
}
