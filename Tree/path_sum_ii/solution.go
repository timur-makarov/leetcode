package path_sum_ii

import (
	"github.com/timur-makarov/leetcode/Implementations/tree"
	"slices"
)

type TreeNode = tree.TreeNode

func RunSolution(root *TreeNode, targetSum int) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	traverse_recursive(root, targetSum, &[]int{}, &result)

	return result
}

func traverse_recursive(node *TreeNode, targetSum int, cPath *[]int, successfulPaths *[][]int) {
	*cPath = append(*cPath, node.Val)

	if node.Val == targetSum && node.Left == nil && node.Right == nil {
		*successfulPaths = append(*successfulPaths, slices.Clone(*cPath))
	}

	if node.Left != nil {
		traverse_recursive(node.Left, targetSum-node.Val, cPath, successfulPaths)
	}

	if node.Right != nil {
		traverse_recursive(node.Right, targetSum-node.Val, cPath, successfulPaths)
	}

	*cPath = (*cPath)[0 : len(*cPath)-1]
}
