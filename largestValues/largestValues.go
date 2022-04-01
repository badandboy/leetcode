package largestValues

import (
	"math"
)

/**
剑指 Offer II 044. 二叉树每层的最大值
官方链接：https://leetcode-cn.com/problems/hPov7L/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func largestValues(root *TreeNode) []int {
	res := map[int]int{}
	var nums []int

	if root == nil {
		return nums
	}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node != nil {
			_, ok := res[depth]
			if !ok {
				res[depth] = node.Val
				nums = append(nums, node.Val)
			}
			res[depth] = int(math.Max(float64(node.Val), float64(res[depth])))
			nums[depth] = res[depth]
			dfs(node.Left, depth+1)
			dfs(node.Right, depth+1)
		}
	}

	dfs(root, 0)

	return nums


}
