package all_nodes_distance_k_in_binary_tree

import "github.com/timur-makarov/leetcode/Implementations/tree"

type TreeNode = tree.TreeNode

func RunSolution(root *TreeNode, target *TreeNode, k int) []int {
	var result []int

	if root == nil {
		return result
	}

	var queue []*TreeNode

	cur := root
	queue = append(queue, cur)
	relations := make(map[int]*TreeNode)

	for len(queue) != 0 {
		levelSize := len(queue)

		for range levelSize {
			cur = queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
				relations[cur.Left.Val] = cur
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
				relations[cur.Right.Val] = cur
			}
		}
	}

	checkedRelations := make(map[int]bool)
	queue = append(queue, target)

	for len(queue) != 0 && k > 0 {
		levelSize := len(queue)

		for range levelSize {
			cur = queue[0]
			queue = queue[1:]
			checkedRelations[cur.Val] = true

			if cur.Left != nil {
				if _, ok := checkedRelations[cur.Left.Val]; !ok {
					queue = append(queue, cur.Left)
				}
			}

			if cur.Right != nil {
				if _, ok := checkedRelations[cur.Right.Val]; !ok {
					queue = append(queue, cur.Right)
				}
			}

			if _, ok := relations[cur.Val]; ok {
				if _, ok := checkedRelations[relations[cur.Val].Val]; !ok {
					queue = append(queue, relations[cur.Val])
				}
			}
		}

		k--
	}

	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		result = append(result, cur.Val)
	}

	return result
}
