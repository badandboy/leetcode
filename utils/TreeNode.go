package utils

import "fmt"

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

func CreateTree(nums []int, i int) *TreeNode {
	if len(nums) <= 0 {
		return &TreeNode{}
	}

	if i >= len(nums) || nums[i] == -1 {
		return nil
	}

	p := &TreeNode{Val: nums[i]}
	p.Left = CreateTree(nums, 2*i+1)
	p.Right = CreateTree(nums, 2*i + 2)

	return p

}

func PrintTree(tree *TreeNode)  {
	if tree != nil {
		fmt.Println(tree.Val)
		PrintTree(tree.Left)
		PrintTree(tree.Right)
	}
}