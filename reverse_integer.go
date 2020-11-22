package leetcode

/*
7. 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

func reverse(x int) int {
	flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	result *= flag
	if result > 1<<31-1 || result < -1<<31 {
		return 0
	}
	return result
}
