/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
var result [][]string
func solveNQueens(n int) [][]string {
	result = [][]string{}

	track := make([][]string, 0, n)
	for i := 0; i < n; i++ {
		row := make([]string, 0, n)
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		track = append(track, row)
	}

	backtrack(n, track, 0)

	return result
}

func backtrack(n int, track [][]string, pick int){
	if pick == n {
		tmp := make([]string, 0, n)
		for _, row := range track {
			str := ""
			for _, v := range row {
				str += v
			}

			tmp = append(tmp, str)
		}

		result = append(result, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if !isValid(track, pick, i) {
			continue
		}

		track[pick][i] = "Q"
		backtrack(n, track, pick + 1)
		track[pick][i] = "."
	}
}

func isValid(track [][]string, r, l int) bool {
	n := len(track)

	// 同一列
	for i := 0; i < n; i++ {
		if track[i][l] == "Q" {
			return false
		}
	}

	// 右上角
	for i, j := r - 1, l + 1; i >= 0 && j < n; i,j=i-1, j+1 {
		if track[i][j] == "Q" {
			return false
		}
	}

	// 左上角
	for i, j := r - 1, l - 1; i >= 0 && j >= 0; i,j=i-1,j-1 {
		if track[i][j] == "Q" {
			return false
		}
	}

	return true
}
// @lc code=end

