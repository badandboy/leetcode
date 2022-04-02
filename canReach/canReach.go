package canReach

/**
这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 i 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。

请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。

注意，不管是什么情况下，你都无法跳到数组之外。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func CanReach(nums []int, start int) bool {
	existMap := make(map[int]int, len(nums))
	var dfs func(nums []int, i int) bool
	dfs = func(nums []int, i int) bool {
		if i < 0 || i >= len(nums) {
			return false
		}

		if _, ok := existMap[i]; ok {
			return false
		}

		if nums[i] == 0 {
			return true
		}

		existMap[i] = -1
		return dfs(nums, i-nums[i]) || dfs(nums, i+nums[i])
	}

	return dfs(nums, start)

}
