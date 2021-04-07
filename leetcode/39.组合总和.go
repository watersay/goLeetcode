/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
var res [][]int
func combinationSum(candidates []int, target int) [][]int {
	// 回溯
	res = make([][]int, 0)
	helper(candidates, target, []int{})

	return res
}

func helper(candidates []int, target int, path []int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for _, v := range candidates {
		if v > target {
			continue
		}

		if len(path) > 0 && v < path[len(path)-1] { // 保证答案有序 不重复
			continue  
		}

		path = append(path, v)
		target -= v
		helper(candidates, target, path)
		path = path[:len(path)-1]
		target += v
	}
}
// @lc code=end

