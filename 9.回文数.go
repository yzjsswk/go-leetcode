package goleetcode

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

/*
 * 模拟
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	dig := []int{}
	for x != 0 {
		dig = append(dig, x%10)
		x /= 10
	}
	l := len(dig)
	i, j := 0, l-1
	for i <= j {
		if dig[i] != dig[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

// @lc code=end
