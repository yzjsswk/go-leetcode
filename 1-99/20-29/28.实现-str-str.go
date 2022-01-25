package goleetcode

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

/*
 * kmp
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	nxt := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j != 0 && needle[i] != needle[j] {
			j = nxt[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		nxt[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = nxt[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
			// j = nxt[j-1]
		}
	}
	return -1
}

// @lc code=end
