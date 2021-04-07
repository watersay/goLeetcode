/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}

	l, r := 0, len(s) - 1

	for l < r {
		vi, vj := rune(s[l]), rune(s[r])
		if !isValid(vi) {
			l++
			continue
		} 
		if !isValid(vj) {
			r--
			continue
		}

		if unicode.ToUpper(vi) != unicode.ToUpper(vj) {
			return false
		}

		l++
		r--
	}

	return true
}
// @lc code=end

