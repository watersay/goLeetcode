/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	// dp 数组定义 [i,j] 区间内的最长回文

	n := len(s)
	dp := make([][]int, n)

	// base case
	for i := 0; i < n; i++ {
		dp[i] =  make([]int, n + 1)
		dp[i][i] = 1
	} 

	for i := n - 2; i >= 0; i-- { // 保证i+1存在
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

