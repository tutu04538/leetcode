/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ptr1, ptr2 := list1, list2
	if ptr1 == nil {
		return ptr2
	}
	if ptr2 == nil {
		return ptr1
	}
	var start *ListNode
	if ptr1.Val > ptr2.Val {
		start = ptr2
		ptr2 = ptr2.Next
	} else {
		start = ptr1
		ptr1 = ptr1.Next
	}
	current := start
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			current.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			current.Next = ptr2
			ptr2 = ptr2.Next
		}
		current = current.Next
	}
	if ptr1 == nil {
		current.Next = ptr2
	} else {
		current.Next = ptr1
	}

	return start
}

// @lc code=end
