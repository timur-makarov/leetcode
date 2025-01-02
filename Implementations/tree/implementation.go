package tree

import (
	"errors"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func RunTreeFunction() {
	tree := &Tree{}

	_ = tree.Insert(5)
	_ = tree.Insert(4)
	_ = tree.Insert(3)
	_ = tree.Insert(8)
	_ = tree.Insert(7)
	_ = tree.Insert(6)
	_ = tree.Insert(1)
	_ = tree.Insert(2)
	_ = tree.Insert(9)

	log.Println("Does it contain 9?: ", tree.Contains(9))
	log.Println("Does it contain 6?: ", tree.ContainsRecursive(6, tree.Root))

	tree.Root, _ = tree.InsertRecursive(12, tree.Root)
	log.Println("Does it contain 12?: ", tree.ContainsRecursive(12, tree.Root))

	tree.Root = tree.RemoveRecursive(12, tree.Root)
	log.Println("Does it contain 12?: ", tree.ContainsRecursive(12, tree.Root))

	tree.PrintBF()
	tree.PrintDFPre()
	tree.PrintDFPost()
	tree.PrintDFIn()
}

func (t *Tree) Insert(val int) error {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
		return nil
	}

	cur := t.Root

	for {
		if cur.Val == val {
			return errors.New("values of nodes should not repeat")
		}

		if cur.Val > val {
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val}
				break
			}
			cur = cur.Left
			continue
		}

		if cur.Val < val {
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}
				break
			}
			cur = cur.Right
			continue
		}
	}

	return nil
}

func (t *Tree) Contains(val int) bool {
	if t.Root == nil {
		return false
	}

	cur := t.Root

	for cur != nil {
		if cur.Val == val {
			return true
		}

		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	return false
}

func (t *Tree) ContainsRecursive(val int, node *TreeNode) bool {
	if node == nil {
		return false
	}

	if node.Val == val {
		return true
	}

	if node.Val > val {
		return t.ContainsRecursive(val, node.Left)
	} else {
		return t.ContainsRecursive(val, node.Right)
	}
}

func (t *Tree) InsertRecursive(val int, node *TreeNode) (*TreeNode, error) {
	var err error

	if node == nil {
		return &TreeNode{Val: val}, nil
	}

	if node.Val == val {
		return nil, errors.New("values of nodes should not repeat")
	}

	if node.Val > val {
		node.Left, err = t.InsertRecursive(val, node.Left)
	} else {
		node.Right, err = t.InsertRecursive(val, node.Right)
	}

	return node, err
}

func (t *Tree) RemoveRecursive(val int, node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}

	if node.Val == val {
		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		minVal := t.MinimalValue(node.Right)
		node.Val = minVal
		node.Right = t.RemoveRecursive(minVal, node.Right)
		return node
	}

	if node.Val > val {
		node.Left = t.RemoveRecursive(val, node.Left)
	} else {
		node.Right = t.RemoveRecursive(val, node.Right)
	}

	return node
}

func (t *Tree) MinimalValue(node *TreeNode) int {
	if node == nil {
		return 0
	}

	for node.Left != nil {
		node = node.Left
	}

	return node.Val
}

func (t *Tree) PrintBF() {
	var queue []*TreeNode
	var results []int

	cur := t.Root

	queue = append(queue, cur)

	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		results = append(results, cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	log.Println(results)
}

func (t *Tree) PrintDFPre() {
	log.Println(t.TraversePre(t.Root))
}

func (t *Tree) PrintDFPost() {
	log.Println(t.TraversePost(t.Root))
}

func (t *Tree) PrintDFIn() {
	log.Println(t.TraverseIn(t.Root))
}

func (t *Tree) TraversePre(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	results := []int{node.Val}

	if node.Left != nil {
		results = append(results, t.TraversePre(node.Left)...)
	}

	if node.Right != nil {
		results = append(results, t.TraversePre(node.Right)...)
	}

	return results
}

func (t *Tree) TraversePost(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	var results []int

	if node.Left != nil {
		results = append(results, t.TraversePost(node.Left)...)
	}

	if node.Right != nil {
		results = append(results, t.TraversePost(node.Right)...)
	}

	results = append(results, node.Val)

	return results
}

func (t *Tree) TraverseIn(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	var results []int

	if node.Left != nil {
		results = append(results, t.TraverseIn(node.Left)...)
	}

	results = append(results, node.Val)

	if node.Right != nil {
		results = append(results, t.TraverseIn(node.Right)...)
	}

	return results
}
