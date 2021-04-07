/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	// 回溯算法 用于取得所有答案
	// 只返回个数可以使用动态规划 背包问题

	// 为什么必须先coins循环？

	dp := make([]int, amount + 1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 1; i <= amount; i ++ {
			if i - coin < 0 { // 凑不出
				continue
			}
			dp[i] += dp[i - coin]
		}
	}

	return dp[amount]
}

// @lc code=end

