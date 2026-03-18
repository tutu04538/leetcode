/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 随机链表的复制
 */
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)

	NewHead := &Node{Val: head.Val, Next: nil, Random: nil}
	m[head] = NewHead
	current := head.Next
	prev := NewHead
	for current != nil {
		NewNode := &Node{Val: current.Val, Next: nil, Random: nil}
		prev.Next = NewNode
		m[current] = NewNode
		prev = NewNode
		current = current.Next
	}

	current = head
	NewCurrent := NewHead

	for current != nil {
		if current.Random == nil {
			NewCurrent.Random = nil
		} else {
			NewCurrent.Random = m[current.Random]
		}
		current = current.Next
		NewCurrent = NewCurrent.Next
	}

	return NewHead

}

// @lc code=end
