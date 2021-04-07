/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x % 10 == 0 {
		return false
	}

	var y int // 反转之后的整数
	for x > y { // 反转一半
		r := x % 10
		x = x / 10
		y = y * 10 + r

		if x == y || x / 10 == y { // x==y 对应len偶数 x / 10 == y 对应len奇数
			return true
		}
	}

	return false
}
// @lc code=end

