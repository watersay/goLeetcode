/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
var res [][]int
func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)

	backtrace(n, k, 1, path)

	return res
}

func backtrace(n int, k int, index int, path []int){
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := index; i <= n; i++ {
		// if i <= index {
		// 	continue
		// }

		path = append(path, i)
		backtrace(n, k, i + 1, path)
		path = path[:len(path)-1]
	}
}
// @lc code=end

