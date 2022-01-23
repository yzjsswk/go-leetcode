package goleetcode

import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

/*
 * 先排序，固定其中两个数，剩下的从一个最小的和最大的开始考虑，由于有序性，
 * 每当和偏小时一定是小的那个变大，每当和偏大时一定是大的那个变小
 * 因为答案不能重复，枚举时要跳过重复的一段到下一个不同的数字
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	ans := [][]int{}
	for a := 0; a < l; {
		for b := a + 1; b < l; {
			for c, d := b+1, l-1; c < d; {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
				}
				if sum < target {
					c++
					for c < d && nums[c] == nums[c-1] {
						c++
					}
				} else {
					d--
					for d > c && nums[d] == nums[d+1] {
						d--
					}
				}
			}
			b++
			for b < l && nums[b] == nums[b-1] {
				b++
			}
		}
		a++
		for a < l && nums[a] == nums[a-1] {
			a++
		}
	}
	return ans
}

// @lc code=end
