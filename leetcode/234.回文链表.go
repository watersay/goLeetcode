/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 var left *ListNode
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	left = head
	return traverse(head)
}

func traverse(head *ListNode) bool {
	if head == nil {
		return true
	}

	res := traverse(head.Next)
	if res && head.Val == left.Val {
		left = left.Next
		return true
	}

	return false
}
// @lc code=end

