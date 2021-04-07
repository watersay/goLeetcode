/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// 考虑每一个柱子能接多少
	// 是由 min(leftMax, rightMax) - cur 得出的
	// 简单实现可以先计算leftMaxMap 和 rightMaxMap 然后遍历就好了

	// 这里使用双指针简化写法
	if len(height) <= 2 {
		return 0
	}

	res := 0

	l, r := 0, len(height) - 1
	leftMax, rightMax := height[0], height[len(height) - 1]

	for l <= r {
		leftMax = max(height[l], leftMax)
		rightMax = max(height[r], rightMax)

		if leftMax > rightMax {
			res += rightMax - height[r]
			r--
		} else {
			res += leftMax - height[l]
			l++
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

