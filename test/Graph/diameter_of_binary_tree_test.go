package Graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Graph/diameter_of_binary_tree"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	var tr Tree

	_ = tr.Insert(9)
	_ = tr.Insert(3)
	_ = tr.Insert(20)
	_ = tr.Insert(15)
	_ = tr.Insert(10)
	_ = tr.Insert(25)
	_ = tr.Insert(1)
	_ = tr.Insert(4)

	answer := 6
	result := diameter_of_binary_tree.RunSolution(tr.Root)
	assert.Equal(t, answer, result)
}
