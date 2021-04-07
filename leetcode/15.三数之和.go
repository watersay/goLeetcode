/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 { // 最小的大于0就不用找了 
			break
		}

		if i > 0 && nums[i] == nums[i-1] { // 轮询中让过重复的
			continue
		}

		l, r := i + 1, len(nums) - 1 // 双指针
		for l < r {
			if l > i + 1 && nums[l] == nums[l-1] { // 防止重复
				l++
				continue
			}
			if r < len(nums) - 1 && nums[r] == nums[r+1] { // 防止重复
				r--
				continue
			}

			if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else if nums[i] + nums[l] + nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]}) // 命中之后别忘了++ 和 --
				r--
				l++
			}
		}
		
	}
	return res
}
// @lc code=end

