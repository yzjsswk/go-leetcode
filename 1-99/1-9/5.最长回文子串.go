package goleetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

/*
 * 枚举中心，模拟向两边拓展计算每个中心的答案
 * 中心有一个的如bab，有两个的如abba，要分别考虑
 */

// @lc code=start
func longestPalindrome(s string) string {
	max_len := 1
	ans_start := 0
	s_len := len(s)
	for c := 0; c < s_len; c++ {
		cur_len := 1
		cur_start := c
		i, j := c-1, c+1
		for i >= 0 && i < s_len && j >= 0 && j < s_len && s[i] == s[j] {
			cur_len += 2
			cur_start = i
			i, j = i-1, j+1
		}
		if cur_len > max_len {
			max_len = cur_len
			ans_start = cur_start
		}

	}
	for c1, c2 := 0, 1; c2 < s_len; c1, c2 = c1+1, c2+1 {
		cur_len := 0
		cur_start := 0
		i, j := c1, c2
		for i >= 0 && i < s_len && j >= 0 && j < s_len && s[i] == s[j] {
			cur_len += 2
			cur_start = i
			i, j = i-1, j+1
		}
		if cur_len > max_len {
			max_len = cur_len
			ans_start = cur_start
		}
	}
	return s[ans_start : ans_start+max_len]
}

// @lc code=end
