package triangleNumber

import "sort"

func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	len1 := len(nums)
	num := 0
	for i := len1 - 1; i >= 2; i-- {
		j := 0
		k := i - 1
		for j < k {
			if nums[i]-nums[k] < nums[j] {
				num += k - j
				k -= 1
			} else {
				j += 1
			}
		}

	}

	return num
}
