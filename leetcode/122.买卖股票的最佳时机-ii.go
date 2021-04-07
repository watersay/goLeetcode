/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 动态规划 变量 当前第几天、 是否持有股票、 对应利润
	// dp := map[int]map[int]int{}
	// dp[0] = make(map[int]int)
	// dp[0][1] = -prices[0]
	// dp[0][0] = 0

	// for i := 1; i < len(prices); i++ {
	// 	dp[i] = map[int]int{} // 初始化
	// 	dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	// 	dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	// }

	// return dp[len(prices) - 1][0]

	// 简写如下 因为看起来只需要两个变量代表dp数组
	// dp_i_1, dp_i_0 := -prices[0], 0
	// for i := 1; i < len(prices); i++ {
	// 	dp_i_1 = max(dp_i_1, dp_i_0 - prices[i])
	// 	dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
	// }
	// return dp_i_0

	// 安规律最简解法
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i - 1] {
			res += prices[i] - prices[i-1]
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

