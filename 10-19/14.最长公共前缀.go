package goleetcode

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

/*
 * 模拟
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	r := 0
	for {
		c := byte(' ')
		for _, s := range strs {
			if len(s) <= r {
				return strs[0][:r]
			}
			if c == ' ' {
				c = s[r]
			} else if c != s[r] {
				return strs[0][:r]
			}
		}
		r++
	}
}

// @lc code=end
