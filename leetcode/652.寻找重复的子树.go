/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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
 var res []*TreeNode
 var dist map[string]int
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res = []*TreeNode{}

	if root == nil {
		return res
	}

	// 每个对比子树的根节点需要做什么？
	// 需要知道自己的整棵树 长什么样子。 然后对比的时候用map，因为不能重复 所以map记录出现次数
	// 知道自己长啥样 后续遍历
	
	dist = make(map[string]int)

	helper(root)
	return res
}

func helper(root *TreeNode) string {
	if root == nil {
		return "#" // 标识结束
	}

	left := helper(root.Left)
	right := helper(root.Right)
	// 这里注意分隔符！ 不然会出现21 和 2 和 1的冲突
	desc := left + "," + right + "," + strconv.Itoa(root.Val)

	freq := dist[desc]
	if freq == 1 {
		res = append(res, root)
	}
	dist[desc] += 1

	return desc
}
// @lc code=end

