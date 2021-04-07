/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start

// 回溯公示
// result = []
// func backtrack(路径，选择列表) {
// 	if 满足结束条件 {
// 		result.add(路径)
// 	}
// 	return

// 	for 选择 in 选择列表 {
// 		做选择
// 		backtrack(路径，选择列表)
// 		撤销选择
// 	}
// }

var res [][]int
func permute(nums []int) [][]int {
	res = [][]int{} // 清空全局数组 leetcode多次执行全局变量不消失

	if len(nums) == 0 {
		return res
	}

	pathNums := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	backtrack(nums, pathNums, used)

	return res
}

func backtrack(nums []int, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		// copy path 因为后面还会修改回溯
		tmp := make([]int, len(pathNums))
		copy(tmp, pathNums)
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true // 标记使用过
			pathNums = append(pathNums, nums[i]) // 做选择
			backtrack(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums) - 1] // 撤销选择
			used[i] = false
		}
	}
}
// @lc code=end

