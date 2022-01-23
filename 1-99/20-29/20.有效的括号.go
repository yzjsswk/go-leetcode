package goleetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

/*
 * 用一个栈，遇到左括号直接入栈，遇到右括号如果栈顶不是对应左括号就不合法，否则出栈
 */

// @lc code=start
func isValid(s string) bool {
	stack := []int{}
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		}
		if s[i] == '[' {
			stack = append(stack, 1)
		}
		if s[i] == '{' {
			stack = append(stack, 2)
		}
		if s[i] == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if s[i] == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != 1 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if s[i] == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != 2 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// @lc code=end
