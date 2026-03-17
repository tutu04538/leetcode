/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	var prev *ListNode
	res := current.Next
	for {
		if current == nil {
			break
		}
		next1 := current.Next
		if next1 == nil {
			break
		}
		next2 := next1.Next
		next1.Next = current
		current.Next = next2
		if prev != nil {
			prev.Next = next1
		}
		prev = current
		current = next2
	}

	return res
}

// @lc code=end
