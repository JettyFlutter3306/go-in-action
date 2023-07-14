package leetcode

import "math"

/*
NO.9  简单  回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	temp := 0
	y := x

	for y > 0 {
		a := y % 10
		temp = temp*10 + a
		y /= 10
	}

	return temp == x
}

/*
NO.7  中等  整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。
*/
func reverse(x int) (res int) {
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		temp := x % 10
		res = res*10 + temp
		x /= 10
	}
	return
}
