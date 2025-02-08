package binary_tree_level_order_traversal

import (
	"github.com/timur-makarov/leetcode/Implementations/tree"
)

type TreeNode = tree.TreeNode

func RunSolution(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	var queue []*TreeNode

	cur := root
	queue = append(queue, cur)
	result = append(result, []int{cur.Val})

	for len(queue) != 0 {
		levelSize := len(queue)
		res := make([]int, 0, levelSize/2)

		for range levelSize {
			cur = queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
				res = append(res, cur.Left.Val)

			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
				res = append(res, cur.Right.Val)
			}
		}

		if len(res) > 0 {
			result = append(result, res)
		}
	}

	reverse(result)
	return result
}

func reverse(arr [][]int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
