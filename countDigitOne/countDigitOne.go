package countDigitOne

/**
233. 数字 1 的个数
给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
官方链接：https://leetcode-cn.com/problems/number-of-digit-one/
*/

func CountDigitOne(n int) int {
	num := 0
	for num/10 > 0 {
		num += num / 10
	}
	return num

}
