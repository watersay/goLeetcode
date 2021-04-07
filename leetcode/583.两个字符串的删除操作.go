/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// 此题就是求两字符串的 最长公共子序列长度 然后想减

	maxCharNum := helper(word1, word2)

	return len(word1) + len(word2) - (2*maxCharNum)
}

func helper(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)

	dp[0] = make([]int, n + 1)

	for i := 1 ; i <=m ; i++ {
		dp[i] = make([]int, n+1)

		for j := 1; j <= n ; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

