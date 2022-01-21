package goleetcode

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

/*
 * 找规律，第1行的相邻两个元素之间如果是下三角，下标之差为k1=2*numrows-2，
 * 如果是上三角，下标之差为k2=0，接下来每增加一行，k1-=2，k2+=2
 *
 * [1] numsRows==1时，k1=k2=0，下面的逻辑会陷入死循环，因此单独考虑
 * [2] 如("AB"，4)，row >= l会导致溢出
 */

// @lc code=start

func convert(s string, numRows int) string {
	if numRows == 1 { // [1]
		return s
	}
	ans := []byte{}
	l := len(s)
	for row, k1, k2 := 0, 2*numRows-2, 0; row < numRows && row < l; row++ { // [2] row < l
		i := row
		ans = append(ans, s[i])
		for i < l {
			i += k1
			if i < l && k1 != 0 {
				ans = append(ans, s[i])
			}
			i += k2
			if i < l && k2 != 0 {
				ans = append(ans, s[i])
			}
		}
		k1, k2 = k1-2, k2+2
	}
	return string(ans)
}

// @lc code=end
