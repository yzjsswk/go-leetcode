package goleetcode

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

/*
 * 模拟
 */

// @lc code=start
func intToRoman(num int) string {
	ans := ""
	for num != 0 {
		switch {
		case num >= 1000:
			num -= 1000
			ans += "M"
		case num >= 900:
			num -= 900
			ans += "CM"
		case num >= 500:
			num -= 500
			ans += "D"
		case num >= 400:
			num -= 400
			ans += "CD"
		case num >= 100:
			num -= 100
			ans += "C"
		case num >= 90:
			num -= 90
			ans += "XC"
		case num >= 50:
			num -= 50
			ans += "L"
		case num >= 40:
			num -= 40
			ans += "XL"
		case num >= 10:
			num -= 10
			ans += "X"
		case num >= 9:
			num -= 9
			ans += "IX"
		case num >= 5:
			num -= 5
			ans += "V"
		case num >= 4:
			num -= 4
			ans += "IV"
		default:
			num -= 1
			ans += "I"
		}
	}
	return ans
}

// @lc code=end
