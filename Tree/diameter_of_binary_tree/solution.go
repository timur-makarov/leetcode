package diameter_of_binary_tree

import (
	"github.com/timur-makarov/leetcode/Implementations/tree"
)

type TreeNode = tree.TreeNode

var result int

func RunSolution(root *TreeNode) int {
	result = 0
	traverse_recursive(root)
	return result
}

func traverse_recursive(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := traverse_recursive(node.Left)
	rightHeight := traverse_recursive(node.Right)

	result = max(result, leftHeight+rightHeight)

	return max(leftHeight, rightHeight) + 1
}
