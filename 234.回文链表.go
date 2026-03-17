/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
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

var frontNode *ListNode

func recursion(currentNode *ListNode) bool {
	if currentNode == nil {
		return true
	}
	if !recursion(currentNode.Next) {
		return false
	}
	if currentNode.Val != frontNode.Val {
		return false
	}
	frontNode = frontNode.Next
	return true
}

func isPalindrome(head *ListNode) bool {
	// s := []int{}
	// for head != nil {
	// 	s = append(s, head.Val)
	// 	head = head.Next
	// }
	// for i := range len(s) / 2 {
	// 	if s[i] == s[len(s)-1-i] {
	// 		continue
	// 	}
	// 	return false
	// }

	// return true
	frontNode = head
	return recursion(head)
}

// @lc code=end
