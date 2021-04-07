/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	fast ,slow := 0, 0

	// 先删除0元素
	for fast < len(nums) {
		if nums[fast] == 0 {
			fast++
			continue
		}

		nums[slow] = nums[fast]
		slow++
		fast++
	}

	// 再标记之后的元素都为0  实现交换
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
// @lc code=end

