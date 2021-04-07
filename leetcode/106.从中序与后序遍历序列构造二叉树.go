/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	val := postorder[len(postorder) - 1]
	var index int // 中序根节点坐标 同时也是左子树长度
	for k, v := range inorder {
		if v == val {
			index = k
		}
	}

	root := &TreeNode{Val: val}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}
// @lc code=end

