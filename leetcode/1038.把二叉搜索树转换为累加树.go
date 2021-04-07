/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 把二叉搜索树转换为累加树
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
var sum int
func bstToGst(root *TreeNode) *TreeNode {
	// 二叉搜索树的性质主要是中序遍历
	// 题目是所有值都是比他大的值和自己的和， 所以以从大到小遍历然后累加即可
	
	sum = 0

	traverse(root)
	return root
}

func traverse(root *TreeNode){
	if root == nil {
		return
	}

	traverse(root.Right)
	sum += root.Val
	root.Val = sum

	traverse(root.Left)
}
// @lc code=end

