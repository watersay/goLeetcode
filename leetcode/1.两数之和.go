/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	dist := make(map[int]int)

	for i, v := range nums {
		j, ok := dist[v]
		if ok {
			return []int{i, j}
		}

		dist[target - v] = i
	}

	return []int{}
}
// @lc code=end

