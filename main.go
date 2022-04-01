package main

import "leetcode/utils"

func main() {
	nums := []int{1,2,3,4,5,6,7}
	tree := utils.CreateTree(nums, 0)
	utils.PrintTree(tree)
}