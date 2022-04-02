package main

import (
	"fmt"
	"leetcode/canReach"
)

func main() {
	nums := []int{3, 0, 2, 1, 2}
	start := 2
	fmt.Println(canReach.CanReach(nums, start))
}
