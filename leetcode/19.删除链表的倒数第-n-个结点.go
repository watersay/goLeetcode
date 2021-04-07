/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	var prev *ListNode
	for fast != nil {
		prev = slow

		fast = fast.Next
		slow = slow.Next
	}

	if prev == nil { // bad case
		return head.Next
	}

	prev.Next = slow.Next
	return head
}
// @lc code=end

