/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 迭代实现
	// var prev *ListNode
	// for head!= nil {
	// 	next := head.Next // 保留下个节点
	// 	head.Next = prev // 更改cur指针
	// 	prev, head = head, next // 递进前置节点 和 当前节点
	// }
	// return prev

	// 递归实现
	if head.Next == nil { // 保证返回的节点是last
		return head
	}
	last := reverseList(head.Next) // 返回的头节点一定是last
	head.Next.Next = head
	head.Next = nil // 递归索引树已经建立 head的指针没有其他作用 保证尾巴一定是nil

	return last
}
// @lc code=end

