package goleetcode

import "math"

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

/*
 * 模拟减法，可以用二分试商优化
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	f := dividend ^ divisor
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	ans := 0
	for dividend >= divisor {
		k, d := 1, divisor
		for d < dividend>>1 {
			k <<= 1
			d <<= 1
		}
		dividend -= d
		ans += k
	}
	if f < 0 {
		ans = -ans
	}
	return ans
}

// @lc code=end
