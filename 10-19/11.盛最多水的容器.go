package goleetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

/*
 * 贪心，当[l, r]计算完成，如果把l，r中较大的往里移动一格，区间长度和min(l,r)都只能减少，
 * 答案不可能变大，因此只要考虑将l，r中较小的往里移动一格的情况
 */

// @lc code=start
func maxArea(height []int) int {
	l := len(height)
	i, j := 0, l-1
	ans := 0
	for i < j {
		if height[i] > height[j] {
			ans = max(ans, height[j]*(j-i))
			j--
		} else {
			ans = max(ans, height[i]*(j-i))
			i++
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
