/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 归并排序
func sortList(head *ListNode) *ListNode { // 函数定义为 返回排好序的序列
	if head == nil || head.Next == nil { // 递归的出口
		return head
	}

	// 找出中点
	var preSlow *ListNode  
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		preSlow = slow
		fast = fast.Next.Next
		slow = slow.Next
	}

	preSlow.Next = nil // 断开前链表

	l1 := sortList(head) // 左右排好序的序列
	l2 := sortList(slow) 

	m := mergeList(l1, l2) // 合并两个有序链表
	return m
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res *ListNode
	if l1.Val < l2.Val {
		res = l1
		res.Next = mergeList(l1.Next, l2)
	} else {
		res = l2
		res.Next = mergeList(l1, l2.Next)
	}
	return res
}
// @lc code=end

