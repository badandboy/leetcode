package minDiffInBST

import (
	"leetcode/utils"
	"sort"
)

func MinDiffInBST(root *utils.TreeNode) int {
	min1 := 1000000
	var dfs func(node *utils.TreeNode)
	var vals []int
	dfs = func(node *utils.TreeNode) {
		vals = append(vals, node.Val)
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)

	sort.Ints(vals)
	for i := 0; i < len(vals)-1; i++ {
		if vals[i+1] == vals[i] {
			return 0
		}
		if vals[i+1]-vals[i] < min1 {
			min1 = vals[i+1] - vals[i]
		}
	}

	return min1
}
