package path_sum_iii

import (
	"github.com/timur-makarov/leetcode/Implementations/tree"
)

type TreeNode = tree.TreeNode

func RunSolution(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return traverse_recursive(root, targetSum, &[]int{})
}

func traverse_recursive(node *TreeNode, targetSum int, cPath *[]int) int {
	count := 0

	*cPath = append(*cPath, node.Val)

	count += countMatchedPaths(*cPath, targetSum)

	if node.Left != nil {
		count += traverse_recursive(node.Left, targetSum, cPath)
	}

	if node.Right != nil {
		count += traverse_recursive(node.Right, targetSum, cPath)
	}

	*cPath = (*cPath)[0 : len(*cPath)-1]
	return count
}

func countMatchedPaths(cPath []int, targetSum int) int {
	sum := 0
	count := 0

	for i := len(cPath) - 1; i >= 0; i-- {
		sum += cPath[i]

		if sum == targetSum {
			count++
		}
	}

	return count
}
