package intersection

import "sort"

func Intersection(nums1 []int, nums2 []int) []int {
	var res []int
	map1 := map[int]bool{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if _, ok := map1[nums1[i]]; !ok {
				res = append(res, nums1[i])
				map1[nums1[i]] = true
			}
			i += 1
		} else if nums1[i] > nums2[j] {
			j += 1
		} else {
			i += 1
		}
	}

	return res
}
