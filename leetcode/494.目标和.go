/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	// 回溯算法很容易想到 复杂度相当于满二叉树遍历 2^n 可以再叠加备忘录

	// 需求可以转化为
	// sum(A) - sum(B) = target 
	// sum(A) = sum(B) + target
	// sum(A) + sum(A)= sum(B) + target + sum(A)
	// sum(A) = (sum(nums) + target) / 2
	// 背包问题 拾取几个+ 凑成sum

	numsSum := sumSlice(nums)
	//转换称背包问题 如果不符合条件 其实无解 
	// sum 都小于目标 和 永远都凑不出的奇数
	if (numsSum < S || (numsSum + S) % 2 == 1) { 
        return 0
    }

	sum := (sumSlice(nums) + S) / 2

	return helper(nums, sum)
}

func helper(nums []int, sum int) int {
	dp := make([][]int, len(nums) + 1)

	for i := 0; i <= len(nums); i++ { // i 是最多捡取多少个
		dp[i] = make([]int, sum + 1)

		dp[i][0] = 1 // 所有t = 0 的都有[]这种情况 

		if i == 0 { // base case 让过
			continue
		}

		for j := 0; j <= sum; j++ {
			if j - nums[i - 1] >= 0 { // 是否捡起第i个值 对应的是i - 1坐标
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)][sum] // 全部捡起是答案
}

func sumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum+=v
	}
	return sum
}
// @lc code=end

