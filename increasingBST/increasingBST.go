package increasingBST

import "leetcode/utils"

func IncreasingBST(root *utils.TreeNode) *utils.TreeNode {
	dummyNode := &utils.TreeNode{}
	curr := dummyNode
	//中序遍历，左-》上-》右
	var vals []int
	var inorder func(node *utils.TreeNode)
	inorder = func(node *utils.TreeNode) {
		if node.Left != nil {
			inorder(node.Left)
		}

		vals = append(vals, node.Val)
		if node.Right != nil {
			inorder(node.Right)
		}
	}
	inorder(root)

	for _, val := range vals {
		curr.Right = &utils.TreeNode{Val: val}
		curr = curr.Right
	}

	return dummyNode.Right
}
