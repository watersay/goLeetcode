/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil {
		if fast.Val == slow.Val {
			fast = fast.Next
			continue
		}

		slow.Next = fast
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = nil  // 截断后面的
	return head
}
// @lc code=end

