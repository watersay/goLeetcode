/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 迭代解法
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := new(ListNode) // 哨兵节点
	result := preHead  // 确定root

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}

		preHead = preHead.Next
	}

	if l1 == nil {
		preHead.Next = l2
	}
	if l2 == nil {
		preHead.Next = l1
	}

	return result.Next
}

// 递归解法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 只考虑当前节点要做什么 返回当前节点 并衔接next递归
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res *ListNode
	if l1.Val < l2.Val {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	} else {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	}
	return res
}
// @lc code=end

