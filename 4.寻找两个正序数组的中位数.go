package goleetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

/*
 * 假设在归并序列中，中位数是第k个数，它左边有a个在第一个数组中，剩下的在第二个数组中，考虑
 * 对第一个数组二分找a，此时两个数组都被分成了两部分，左+左对应归并序列中中位数左边的那些数，
 * 正确的a应满足：左边的最大值小于等于右边的最小值，否则，当左边的最大值出现在第一个数组中时，
 * 向左二分，当左边的最大值出现在第二个数组中时，向右二分
 *
 * 如：
 *  1 2 3 4
 *  100 200 300
 * 中位数的左边应该有3个数，对第一个数组二分一次：
 *  1 2 | 3 4
 *  100 | 200 300
 * 此时认为中位数左边是1，2，100，这是不对的，因为左边的最大值100大于右边的最小值3，
 * 且这个100出现在第二个数组中，我们接下来希望把这个100弄到右边去，所以第二个数组的分割线
 * 要左移，所以第一个数组的分割线要右移，也就是向右二分：
 * 1 2 3 | 4
 * | 100 200 300
 * 这时候左边的最大值3小于右边的最小值4了，所以左边的1，2，3就是中位数左边的那些数，
 * 中位数就应该是右边的最小值4(偶数情况有所不同)
 *
 * 具体实现时，边界情况比较多，要仔细考虑
 *
 * [1] 数组为空单独考虑
 * [2] 组成左边的数nums1下标到mid，则nums2下标到k-mid-2，为-1时表示全在右边，
 * ok为true表示这种情况是正确的划分方法，且中位数为ans，否则goleft为true时往左
 * 二分，goleft为false时往右二分
 * [3] bound的范围应该是-1到len-1，小于-1说明nums2中分界线太左边了，不可行且
 * 接下来要往右二分，大于len-1说明nums2中分界线太右边了，不可行且接下来要往左二分，
 * nums1的bound由二分得到，不会出现这种情况
 * [4] 按数字总数的奇偶分别计算答案
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	if n == 0 { // [1]
		return float64(nums2[m/2]+nums2[(m+1)/2-1]) / 2
	}
	if m == 0 {
		return float64(nums1[n/2]+nums1[(n+1)/2-1]) / 2
	}
	k := (n + m) / 2
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if ans, ok, goleft := check(nums1, nums2, mid, k-mid-2); ok { // [2]
			return ans
		} else {
			if goleft {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	ans, _, _ := check(nums1, nums2, -1, k-1)
	return ans
}

func check(nums1, nums2 []int, bound1, bound2 int) (float64, bool, bool) {
	n, m := len(nums1), len(nums2)
	l_max, r_min := -1000005, 1000005
	if bound2 < -1 { // [3]
		return 0.0, false, true
	}
	if bound2 >= m {
		return 0.0, false, false
	}
	if bound1 != -1 {
		l_max = max(l_max, nums1[bound1])
	}
	if bound2 != -1 {
		l_max = max(l_max, nums2[bound2])
	}
	if bound1 != n-1 {
		r_min = min(r_min, nums1[bound1+1])
	}
	if bound2 != m-1 {
		r_min = min(r_min, nums2[bound2+1])
	}
	ans := 0.0
	if (n+m)&1 != 0 { // [4]
		ans = float64(r_min)
	} else {
		ans = (float64(l_max) + float64(r_min)) / 2
	}
	return ans, l_max <= r_min, bound1 != -1 && l_max == nums1[bound1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
