package goleetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

/*
 * 先排序，固定其中一个数k，剩下的从一个最小i的和最大的j开始考虑，由于有序性，
 * 每当和偏小时一定是小的那个变大(i++)，每当和偏大时一定是大的那个变小(j--)
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for k := 0; k < l; k++ {
		for i, j := k+1, l-1; i < j; {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum == target {
				return target
			}
			if sum < target {
				i++
			} else {
				j--
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
