/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	
	var root *ListNode // 返回最小的节点作为根节点
	index := -1 // 记录最小节点坐标

	for i, n := range lists {
		if n == nil {
			continue
		}

		if root == nil {
			root = n
			index = i
			continue
		}

		if n.Val < root.Val {
			root = n
			index = i
		}
	}

	if index == -1 {
		return nil
	}

	// 后移被选取的node
	lists[index] = lists[index].Next

	// 递归
	root.Next = mergeKLists(lists)

	return root
}
// @lc code=end

