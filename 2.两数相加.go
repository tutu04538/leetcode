/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	m, n := l1.Val, l2.Val
	carry = (m + n) / 10
	start := &ListNode{Next: nil, Val: (m + n) % 10}
	current := start
	l1, l2 = l1.Next, l2.Next
	for l1 != nil && l2 != nil {
		m, n := l1.Val, l2.Val
		current.Next = &ListNode{Next: nil, Val: (m + n + carry) % 10}
		carry = (m + n + carry) / 10
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 == nil {
		l2 = l1
	}
	for carry > 0 && l2 != nil {
		current.Next = &ListNode{Next: nil, Val: (carry + l2.Val) % 10}
		carry = (carry + l2.Val) / 10
		current = current.Next
		l2 = l2.Next
	}
	if carry > 0 {
		current.Next = &ListNode{Next: nil, Val: carry}
	} else {
		current.Next = l2
	}
	return start

}

// @lc code=end
