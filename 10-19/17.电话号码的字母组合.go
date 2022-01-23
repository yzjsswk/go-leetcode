package goleetcode

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

/*
 * 模拟，搜索
 */

// @lc code=start

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dic := []string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	ans := []string{}
	var dfs func(int, string)
	dfs = func(deep int, cans string) {
		if deep == len(digits) {
			ans = append(ans, cans)
			return
		}
		s := dic[digits[deep]-48]
		l := len(s)
		for i := 0; i < l; i++ {
			dfs(deep+1, cans+string(s[i]))
		}
		return
	}
	dfs(0, "")
	return ans
}

// @lc code=end
