/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
			continue
		}

		slow++
		nums[slow] = nums[fast]
		fast++
	}

	return slow + 1

	// fast, slow := 0, 0

	// tmp := 2 << 31 - 1
	// for fast < len(nums) {
	// 	if nums[fast] == tmp {
	// 		fast++
	// 		continue
	// 	}

	// 	tmp = nums[fast]
	// 	nums[slow] = nums[fast]
	// 	slow++
	// 	fast++
	// }

	// return slow
}
// @lc code=end

