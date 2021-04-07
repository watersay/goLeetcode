/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
var res [][]int
func subsets(nums []int) [][]int {
	// 这是一个回溯问题
	res = make([][]int, 0)
	vis := make([]bool, len(nums))

	sort.Ints(nums) // 去重复
	backtrace(nums, []int{}, vis)
	return res
}

func backtrace(nums []int, paths []int, vis []bool){
	// 接收所有可能答案
	tmp := make([]int, 0)
	tmp = append(tmp, paths...)
	res = append(res, tmp)

	for i := 0; i < len(nums); i++ {
		// 只能从小到大 取出paths的最后一个元素做比较
		if vis[i] || (len(paths) != 0 && paths[len(paths)-1] > nums[i]) {
			continue
		}

		vis[i] = true
		paths = append(paths, nums[i])
		backtrace(nums, paths, vis)
		vis[i] = false
		paths = paths[:len(paths)-1]
	}
}
// @lc code=end

