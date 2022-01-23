package goleetcode

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

/*
 * dfs，模拟
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ans := []string{}
	var dfs func(int, int, string)
	dfs = func(l, r int, cans string) {
		if l == n && r == n {
			ans = append(ans, cans)
			return
		}
		if l < n {
			dfs(l+1, r, cans+"(")
		}
		if r < l {
			dfs(l, r+1, cans+")")
		}
	}
	dfs(0, 0, "")
	return ans
}

// @lc code=end
