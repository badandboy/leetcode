package findOriginalArray

/**
一个整数数组 original 可以转变成一个 双倍 数组 changed ，转变方式为将 original 中每个元素 值乘以 2 加入数组中，然后将所有元素 随机打乱 。

给你一个数组 changed ，如果 change 是 双倍 数组，那么请你返回 original数组，否则请返回空数组。original 的元素可以以 任意 顺序返回。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-original-array-from-doubled-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func FindOriginalArray(changed []int) []int {
	if len(changed)%2 == 1 {
		return []int{}
	}

}
