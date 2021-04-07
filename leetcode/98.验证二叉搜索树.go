/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

 var tmp int
func isValidBST(root *TreeNode) bool {
	tmp = - 2 << 31
	return traverse(root)
}

func traverse(root *TreeNode) bool { // 方向 限制
	if root == nil {
		return true
	}

	// 这里利用了 中序遍历有序的特性 有序必然是bst
	ok1 := traverse(root.Left)
	if root.Val <= tmp {
		return false
	}
	tmp = root.Val

	ok2 := traverse(root.Right)

	return ok1 && ok2
}
// @lc code=end

