/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2{
		return 0
	}

	// 只卖一次 所以遍历一遍 找到最小值 记录最大差值

	min, maxSale := prices[0], 0

	for _, p := range prices {
		if p < min {
			min = p
		} else {
			if (p - min) > maxSale {
				maxSale = p - min
			}
		}
	}

	return maxSale
}
// @lc code=end

