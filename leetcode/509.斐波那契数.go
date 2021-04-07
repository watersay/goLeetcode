/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n < 2 {
		return n
	}

	prev, cur := 0, 1
	res := 0

	for i := 2; i <= n ; i++ {
		res = prev + cur
		prev, cur = cur, res
	} 
	return res
}
// @lc code=end

