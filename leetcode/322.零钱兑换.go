/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	
	dp := make([]int, amount + 1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = -1

		for _, coin := range coins {
			// dp[i - coin] 凑不出值的子问题必须跳过， 否则会被后续流程破坏值
			if i - coin < 0 || dp[i - coin] == -1 { // 跳过无解子问题
				continue
			}

			if dp[i] == -1 {
				dp[i] = 1 + dp[i - coin]
			} else {
				dp[i] = min(dp[i], 1 + dp[i - coin])
			}
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

