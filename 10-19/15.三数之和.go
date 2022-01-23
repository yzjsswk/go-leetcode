package goleetcode

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
// @lc code=start

/*
 * 先排序，固定其中一个数k，剩下的从一个最小i的和最大的j开始考虑，由于有序性，
 * 每当和偏小时一定是小的那个变大(i++)，每当和偏大时一定是大的那个变小(j--)
 * 因为答案不能重复，每当k，i，j移动时一定要跳过重复的一段到下一个不同的数字
 */

func threeSum(nums []int) [][]int {
	l := len(nums)
	ans := [][]int{}
	sort.Ints(nums)
	for k := 0; k < l; {
		for i, j := k+1, l-1; i < j; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
			if sum < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else {
				j--
				for j > i && nums[j] == nums[j+1] {
					j--
				}
			}
		}
		k++
		for k < l && nums[k] == nums[k-1] {
			k++
		}
	}
	return ans
}

// @lc code=end
