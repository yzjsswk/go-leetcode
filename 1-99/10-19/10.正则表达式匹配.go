package goleetcode

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

/*
 * 这题好烦啊
 */

/*
 "aa"
 "a"
 "aa"
 "a*"
 "ab"
 ".*"
 "aab"
 "c*a*b"
 "mississippi"
 "mis*is*p*."
 "c"
 "c*c"
 "abcabcabcabcd"
 "ab.*abcd"
 "a"
 "ab*"
 "acaabbaccbbacaabbbb"
 "a*.*b*.*a*aa*a*"
*/

// @lc code=start
func isMatch(s string, p string) bool {
	l1, l2 := len(s), len(p)
	i, j := 0, 0
	for {
		if i == l1 {
			for {
				if j == l2 {
					return true
				}
				if j+1 < l2 && p[j+1] == '*' {
					j = j + 2
				} else {
					return false
				}
			}
		}
		if i >= l1 || j >= l2 {
			break
		}
		if j == l2-1 {
			return i == l1-1 && (s[i] == p[j] || p[j] == '.')
		}
		if p[j+1] != '*' {
			if s[i] == p[j] || p[j] == '.' {
				i, j = i+1, j+1
				continue
			} else {
				return false
			}
		}
		if p[j] == '.' {
			for start := i; start <= l1; start++ {
				if isMatch(s[start:], p[j+2:]) {
					return true
				}
			}
			return false
		}
		flag := true
		for start := i; start <= l1; start++ {
			if isMatch(s[start:], p[j+2:]) {
				return true
			}
			if start < l1 && s[start] != p[j] {
				j += 2
				flag = false
				break
			}
		}
		if flag {
			return false
		}
	}
	return false
}

// @lc code=end
