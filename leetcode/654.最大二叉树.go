/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// 先确认根节点要做什么 然后递归下去

	max := 0
	maxIndex := 0
	// 找出最大值 和 最大值坐标
	for i, n := range nums {
		if n > max {
			max = n
			maxIndex = i
		}
	}

	root := &TreeNode{Val:max}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex + 1 :])

	return root
}
// @lc code=end

