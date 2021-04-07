/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	// 先找到相交点， 环长度为k， head到相交点长度为k，让slow回到head再次一步步走 第一个相交为环起点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil { // 如果因为for条件而终止的
		return nil
	}

	// 当前slow 和 fast都在相交点
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
// @lc code=end

