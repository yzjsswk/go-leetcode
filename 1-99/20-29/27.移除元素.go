package goleetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

/*
 * 扫一遍，遇到val就和最后一个交换
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	l := len(nums)
	i := 0
	for i < l {
		if nums[i] == val {
			nums[i], nums[l-1] = nums[l-1], nums[i]
			l--
		} else {
			i++
		}
	}
	return l
}

// @lc code=end
