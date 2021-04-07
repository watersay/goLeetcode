/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 0

	for i := 0; i < len(s); i++ {
		res++

		if i + 1 < len(s) && s[i] == s[i+1] { //子串偶数
			res++

			l, r := i-1, i+2
			for l >= 0 && r < len(s) && s[l] == s[r] {
				res++
				l--
				r++
			} 
		}

		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		} 
	}

	return res
}
// @lc code=end

