package goleetcode

import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

/*
 * 先反转到倒数第二位，和最值的除最后一位比较，分类计算结果
 */

// @lc code=start
func reverse(x int) int {
	ans := 0
	f := 1
	if x < 0 {
		f = -1
		x = -x
	}
	for x >= 10 {
		ans = ans*10 + x%10
		x /= 10
	}
	max_part1 := math.MaxInt32 / 10
	max_part2 := math.MaxInt32 % 10
	if f == -1 {
		max_part2++
	}
	if ans < max_part1 {
		return f*ans*10 + f*x
	}
	if ans > max_part1 {
		return 0
	}
	if x > max_part2 {
		return 0
	} else {
		return f*ans*10 + f*x
	}

}

// @lc code=end
