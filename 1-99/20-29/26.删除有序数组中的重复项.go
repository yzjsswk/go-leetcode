package goleetcode

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

/*
 * 扫一遍，每次遇到不一样的就填到最前面
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l := len(nums)
	i := 0
	k := 0
	for i < l {
		for i < l-1 && nums[i] == nums[i+1] {
			i++
		}
		nums[k] = nums[i]
		k++
		i++
	}
	return k
}

// @lc code=end
