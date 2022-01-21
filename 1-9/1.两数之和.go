package goleetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

/*
 * 哈希记录每个数出现的位置，对于i找target-i是否存在
 * 因为i和target-i总有一个出现时另一个已经被记录，所以可以边找边存，一边遍历完成
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
