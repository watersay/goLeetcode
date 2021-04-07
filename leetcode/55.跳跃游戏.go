/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	// 最值 最大跳到的长度是多少
	// 贪心

	maxJump := 0
	// for i := 0; i <= maxJump; i++ { // 遍历所及 是因为已经可以覆盖到了
	// 	maxJump = max(maxJump, i + nums[i])

	// 	if maxJump >= len(nums) - 1 {
	// 		return true
	// 	}
	// }
	// return false

	for i := 0; i < len(nums); i++ {
		if maxJump < i {
			return false
		}

		maxJump = max(maxJump, i + nums[i])
	}

	return maxJump >= len(nums) - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

