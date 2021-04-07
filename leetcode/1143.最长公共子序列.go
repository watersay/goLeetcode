/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	// 动态规划
	dp := make([][]int, len(text1) + 1) // 数组扩大1 让遍历从1开始 防止-1问题
	// 此时 i j 代表的是位数而不是指针

	dp[0] = make([]int, len(text2) + 1)
	
	for i := 1; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2) + 1)

		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] { // 从已有推未知
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

