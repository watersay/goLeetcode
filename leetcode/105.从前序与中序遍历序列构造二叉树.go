/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 找到根节点 操作root.Left 和 root.Right 递归下去

	val := preorder[0]

	// 找到中旭遍历root位置
	var index int // 同时也是左子树的长度
	for k, v := range inorder {
		if v == val {
			index = k
			break
		}
	}

	root := &TreeNode{Val: val}
	// 画出前序遍历和中序遍历的数据 画一下就出来了   index是中序root位置 同时也是左子树长度
	root.Left = buildTree(preorder[1:index + 1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return root
	
}
// @lc code=end

