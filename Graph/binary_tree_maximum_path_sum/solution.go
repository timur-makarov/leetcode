package binary_tree_maximum_path_sum

import (
	"github.com/timur-makarov/leetcode/Implementations/tree"
	"math"
)

type TreeNode = tree.TreeNode

var result int

func RunSolution(root *TreeNode) int {
	result = math.MinInt
	traverse_recursive(root)
	return result
}

func traverse_recursive(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftSum := traverse_recursive(node.Left)
	rightSum := traverse_recursive(node.Right)

	leftSum = max(leftSum, 0)
	rightSum = max(rightSum, 0)

	result = max(result, leftSum+rightSum+node.Val)

	return max(leftSum, rightSum) + node.Val
}
