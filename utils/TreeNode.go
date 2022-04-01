package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InsertNode(root *TreeNode, val int, left bool) *TreeNode {
	node := &TreeNode{Val: val}
	if left {
		root.Left = node
	} else {
		root.Right = node
	}
	return root
}
