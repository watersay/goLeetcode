/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := []byte(strs[0]) // 选取第一个作为样板
	for i := 1; i < len(strs); i++ { 
		if len(strs[i]) == 0 {
			return ""
		}

		for j := 0; j < len(res); j++ {
			if j >= len(strs[i]) || strs[i][j] != res[j] {
				res = res[:j] // 截断样板
			}
		}
	}

	return string(res)
}
// @lc code=end

