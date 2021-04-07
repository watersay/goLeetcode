/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	// 扩展遍历
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	max := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				num := 0
				dfs(grid, i, j, &num)
				if num > max {
					max = num
				}
			}
		}
	}

	return max
}

func dfs(grid [][]int, i, j int, num *int) {
	*num++
	grid[i][j] = 0 // 标记遍历过
	// 上
	if j - 1 >= 0 && grid[i][j-1] == 1 {
		dfs(grid, i, j-1, num)
	}
	//下
	if j + 1 < len(grid[0]) && grid[i][j+1] == 1 {
		dfs(grid, i, j+1, num)
	}

	if i - 1 >= 0 && grid[i-1][j] == 1 {
		dfs(grid, i - 1, j, num)
	}
	if i + 1 < len(grid) && grid[i+1][j] == 1 {
		dfs(grid, i + 1, j, num)
	}
}
// @lc code=end

