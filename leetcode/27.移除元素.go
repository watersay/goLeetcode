/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0

	for fast < len(nums) {
		if nums[fast] == val {
			fast++
			continue
		}

		nums[slow] = nums[fast]
		slow++
		fast++
	}

	return slow
}
// @lc code=end

