/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	}

	n := len(prices)

	// 如果k > len(prices) 则等于k不限制
	if k > n / 2 {
		return profit(prices)
	}

	// dp状态 第几天、 k、 是否持有
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k + 1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= k; j ++ {
			if i == 0 {
				dp[i][j][1] = -prices[i]
				dp[i][j][0] = 0
				continue
			}
			// 在卖的时候处理k
			// 昨天就持有 和 昨天没持有并买了
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i])
			
			// 昨天没持有 和 昨天持有并卖了
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i])
		}
	}

	return dp[n - 1][k][0]
}

// 不限制交易次数的情况下
func profit(prices []int) int {
	res := 0 
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i - 1] {
			res += prices[i] - prices[i - 1]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

