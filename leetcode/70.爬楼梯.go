/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	pre, cur := 1, 2
	res := 0

	for i := 3; i <= n; i++ {
		res = pre + cur
		pre = cur
		cur = res
	}

	return res
}
// @lc code=end

