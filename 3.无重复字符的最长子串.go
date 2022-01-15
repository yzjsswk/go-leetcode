package goleetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

/*
 * 考虑i和j之间的一段子串，保证i和j之间没有重复的字符
 * 记录每种字符最后出现的位置，当j遇到重复字符时，i往前跳跃使i和j之间没有重复字符
 * [1]  例如abccb，当j遇到第二个c时，i往前跳跃，跳过了第一个b，导致map中b的值未更新，
 *		当遇到第二个b时，i会往回跳，这是不对的，i < p 避免了这种情况
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	ans := 0
	mp := map[byte]int{}
	i, j := 0, 0
	l := len(s)
	for j < l {
		c := s[j]
		if p, ok := mp[c]; ok && i < p { //[1] i < p
			i = p
		}
		j++
		mp[c] = j
		if j-i > ans {
			ans = j - i
		}
	}
	return ans
}

// @lc code=end
