/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		s1 := helper(s, i, i) 
		s2 := helper(s, i, i + 1)

		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}

	return res
}

// 传入指针 以lr为中信的找到最大回文子串
func helper(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	l++
	r-- // 前面多+了 回撤

	return s[l : r + 1] // 左开右闭
}
// @lc code=end

