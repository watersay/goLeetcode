/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root) // 栈

	for len(stack) != 0 {
		res = append(res, []int{})

		tmpStack := make([]*TreeNode, 0) // 创建临时栈
		// 读取这一层的数据 写入res
		for _, node := range stack {
			res[len(res)-1] = append(res[len(res)-1], node.Val)

			if node.Left != nil {
				tmpStack = append(tmpStack, node.Left)
			}
			if node.Right != nil {
				tmpStack = append(tmpStack, node.Right)
			}
		}

		// 替换栈
		stack = tmpStack
	}

	return res
}
// @lc code=end

