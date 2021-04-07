/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	res := ""

	carry := 0 

	for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || carry != 0; i, j = i - 1, j - 1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0') // assic码转换
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		tmp := x + y + carry
		res = strconv.Itoa(tmp % 10) + res
		carry = tmp / 10
	}

	return res
}

// @lc code=end

