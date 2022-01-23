package goleetcode

import "math"

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

/*
 * 模拟
 */

// @lc code=start
func myAtoi(s string) int {
	l := len(s)
	p := 0
	for p < l {
		if s[p] == ' ' {
			p++
		} else {
			break
		}
	}
	if p == l {
		return 0
	}
	f := 1
	if s[p] == '-' {
		f = -1
		p++
	} else if s[p] == '+' {
		p++
	}
	var ans int64
	for p < l {
		if s[p] < '0' || s[p] > '9' {
			break
		}
		ans = ans*10 + int64(s[p]) - 48
		if ans > math.MaxInt32 {
			if f == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		p++
	}
	return f * int(ans)
}

// @lc code=end
