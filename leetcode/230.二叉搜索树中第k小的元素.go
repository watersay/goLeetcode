/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
 var rank int
 var res int
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	rank = 0 // 初始化
	res = 0
	
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}

	traverse(root.Left, k)
	rank++ // rank++ 必须在判断之前 否则会导致递归回去的时候全程覆盖res
	if rank == k {
		res = root.Val
		return
	}

	traverse(root.Right, k)
}
// @lc code=end

