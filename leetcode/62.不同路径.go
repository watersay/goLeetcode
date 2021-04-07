/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m + 1)

	for i :=0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m ; i++ {
		for j := 1; j <= n ; j ++ {
			if i == 1 && j == 1 {
				dp[i][j] = 1 // base case
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
// @lc code=end

