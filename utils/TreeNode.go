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

//func createTree(nums []int) *TreeNode {
//	if len(nums) <= 0 {
//		return &TreeNode{}
//	}
//
//	dummyNode := &TreeNode{}
//	curr := dummyNode
//	for _, val := range nums {
//		if  {
//
//		}
//		curr.Left = &TreeNode{Val: val}
//	}
//}
