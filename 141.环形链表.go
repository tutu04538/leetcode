/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
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
func hasCycle(head *ListNode) bool {
	// m := make(map[*ListNode]struct{})
	// for head != nil {
	// 	if _, ok := m[head]; ok {
	// 		return true
	// 	}
	// 	m[head] = struct{}{}
	// 	head = head.Next
	// }
	// return false
	// Floyd 判圈算法/龟兔赛跑算法
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// @lc code=end
