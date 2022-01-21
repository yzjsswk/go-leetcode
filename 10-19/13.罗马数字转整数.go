package goleetcode

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

/*
 * 模拟
 */

// @lc code=start
func romanToInt(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case 'V':
			if i > 0 && s[i-1] == 'I' {
				ans += 4
				i--
			} else {
				ans += 5
			}
		case 'X':
			if i > 0 && s[i-1] == 'I' {
				ans += 9
				i--
			} else {
				ans += 10
			}
		case 'L':
			if i > 0 && s[i-1] == 'X' {
				ans += 40
				i--
			} else {
				ans += 50
			}
		case 'C':
			if i > 0 && s[i-1] == 'X' {
				ans += 90
				i--
			} else {
				ans += 100
			}
		case 'D':
			if i > 0 && s[i-1] == 'C' {
				ans += 400
				i--
			} else {
				ans += 500
			}
		case 'M':
			if i > 0 && s[i-1] == 'C' {
				ans += 900
				i--
			} else {
				ans += 1000
			}
		default:
			ans += 1
		}
	}
	return ans
}

// @lc code=end
