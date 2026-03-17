/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 在环中相遇，计算环的长度 b
			ring_len := 1
			slow = slow.Next
			for slow != fast {
				slow = slow.Next
				ring_len += 1
			}
			// 先让 ptr1 从 b，再让 ptr1 ptr2 同时走，相遇时即为环的起点
			// 证明，假设非环的长度是 a，ptr1 走 a+b+b，ptr2 走 a+b 时，相遇，此时在环的起点
			ptr1, ptr2 := head, head
			for range ring_len {
				ptr1 = ptr1.Next
			}
			for ptr1 != ptr2 {
				ptr1 = ptr1.Next
				ptr2 = ptr2.Next
			}
			return ptr1

		}
	}

	return nil
}

// @lc code=end
