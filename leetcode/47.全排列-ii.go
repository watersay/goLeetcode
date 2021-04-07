/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
var res [][]int
func permuteUnique(nums []int) [][]int {
	// 去重复 需要先堆nums排序后才好最判断
	sort.Ints(nums)

	res = make([][]int, 0)
	vis := make([]bool, len(nums))

	backtrace(nums, []int{}, vis)

	return res
}

func backtrace(nums []int, paths []int, vis []bool) {
	if len(paths) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, paths)
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 位置被访问过 
		// 或者 前后相等 并且前一个没有被放进去（放入有顺序来保证不重复）
		if vis[i] || (i > 0 && nums[i] == nums[i-1] && !vis[i - 1]) { 
			continue
		}

		// 做选择
		vis[i] = true 
		paths = append(paths, nums[i])
		backtrace(nums, paths, vis)
		// 撤销选择
		vis[i] = false
		paths = paths[:len(paths)-1]
	}
}
// @lc code=end

