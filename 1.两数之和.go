package goleetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	for i, num := range nums {
		if j, ok := mp[target-num]; ok {
			return []int{i, j}
		}
		mp[num] = i
	}
	return nil
}

// @lc code=end
