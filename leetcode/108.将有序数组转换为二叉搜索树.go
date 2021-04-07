/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	// 二茬搜索书的 中序遍历是 有序输出
	if len(nums) == 0 {
		return nil
	}

	var res *TreeNode = &TreeNode{}
	mid := len(nums) / 2

	res.Val = nums[mid]
	res.Left = sortedArrayToBST(nums[:mid])
	res.Right = sortedArrayToBST(nums[mid+1:])

	return res
}
// @lc code=end

