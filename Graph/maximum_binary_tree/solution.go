package maximum_binary_tree

import "github.com/timur-makarov/leetcode/Implementations/tree"

type TreeNode = tree.TreeNode

func RunSolution(nums []int) *TreeNode {
	var stack []*TreeNode

	for _, num := range nums {
		node := &TreeNode{Val: num}

		for len(stack) != 0 && node.Val > stack[len(stack)-1].Val {
			node.Left = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) != 0 {
			stack[len(stack)-1].Right = node
		}

		stack = append(stack, node)
	}

	return stack[0]
}
