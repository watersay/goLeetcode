/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := helper(root.Left)
	right := helper(root.Right)

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right
	return root
}
// @lc code=end

