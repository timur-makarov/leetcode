package kth_smallest_element_in_a_bst

import "github.com/timur-makarov/leetcode/Implementations/tree"

type TreeNode = tree.TreeNode

func RunSolution(root *TreeNode, k int) int {
	result := 0
	count := 0
	traverseIn(root, &count, &result, k)
	return result
}

func traverseIn(node *TreeNode, count, result *int, k int) {
	if node == nil {
		return
	}

	traverseIn(node.Left, count, result, k)

	*count++
	if *count == k {
		*result = node.Val
		return
	}

	traverseIn(node.Right, count, result, k)
}
