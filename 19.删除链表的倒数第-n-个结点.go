/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	first, second := head, head
	for range n - 1 {
		first = first.Next
	}
	if first.Next == nil {
		return head.Next
	}
	var prev *ListNode
	for {
		first = first.Next
		if first == nil {
			prev.Next = second.Next
			return head
		}
		prev = second
		second = second.Next
	}

}

// @lc code=end
