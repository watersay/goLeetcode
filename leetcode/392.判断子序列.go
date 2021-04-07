/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	// 双指针
	i, j := 0, 0
	sum := 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			sum++
		}
		j++
	}

	return sum == len(s)
}
// @lc code=end

