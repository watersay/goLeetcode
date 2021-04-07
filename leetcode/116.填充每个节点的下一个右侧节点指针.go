/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
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
	if root == nil {
		return nil
	}

	// 由root出发可以处理左右 但是无法处理孙子节点。 那就递进一步
	helper(root.Left, root.Right)

	return root
}

func helper(a, b *Node) {
	if a == nil || b == nil {
		return
	}

	a.Next = b

	helper(a.Left, a.Right)
	helper(b.Left, b.Right)

	helper(a.Right, b.Left)
}
// @lc code=end

