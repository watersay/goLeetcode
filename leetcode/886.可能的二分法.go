/*
 * @lc app=leetcode.cn id=886 lang=golang
 *
 * [886] 可能的二分法
 */

// @lc code=start
func possibleBipartition(n int, dislikes [][]int) bool {

	// 整理dislikes 数据为节点映射
	graph := make([][]int, n + 1)
	for _, v := range dislikes {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	colors := make([]int, n+1) // 记录染色 1 和 -1

	for i := 1; i <= n; i++ {
		if colors[i] != 0 { // 如果染过色
			continue
		}

		colors[i] = 1 // 默认指定1
		// bfs
		queue := []int{i}
		for len(queue) != 0 {
			n := queue[0]
			queue = queue[1:]
			color := colors[n]
			for _, v := range graph[n] { // 遍历讨厌队列
				if colors[v] == color { // 如果已经存在颜色并且冲突 返回失败
					return false
				}

				// 处理过的不需要重复入队
				if colors[v] != 0 {
					continue
				}

				// 入队
				queue = append(queue, v)

				// 染色
				if colors[v] == 0 {
					colors[v] = -color 
				} 
			}
		}
	}

	return true
}
// @lc code=end

