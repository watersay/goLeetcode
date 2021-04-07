/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res := ^int(^uint(0) >> 1)
	prev := 0

	// 动态规划 计算必须包含i数的最大值  然后中间的所有答案做max就出结果了
	for i := 0; i < len(nums); i++ {
		prev = max(nums[i], prev + nums[i]) // 要与不要
		res = max(res, prev)
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

