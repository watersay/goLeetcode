/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	// 先把广度优先遍历写下来 再处理细节
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		tmp := queue
		queue = nil 

		for i, n := range tmp {
			if i + 1 <= len(tmp) - 1 {
				n.Next = tmp[i + 1]
			}
			
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return root
}
// @lc code=end

