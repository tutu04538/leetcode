/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recursiveInorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	recursiveInorderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	recursiveInorderTraversal(root.Right, res)
}

type Stack94 struct {
	items []*TreeNode
}

func (s *Stack94) Push(item *TreeNode) {
	s.items = append(s.items, item)
}

func (s *Stack94) Pop() (*TreeNode, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("Stack94 is empty")
	}
	res := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return res, nil
}

func (s *Stack94) Peek() (*TreeNode, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("Stack94 is empty")
	}
	res := s.items[len(s.items)-1]
	return res, nil
}

func (s *Stack94) IsEmpty() bool {
	return len(s.items) == 0
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// res := []int{}
	// recursiveInorderTraversal(root, &res)
	// return res

	// 迭代方式中序遍历
	// 使用栈进行维护，每次 pop 时，说明该子树的左子树已经遍历完毕
	res := []int{}
	Stack94 := Stack94{}
	cur := root
	for cur != nil {
		Stack94.Push(cur)
		cur = cur.Left
	}
	for !Stack94.IsEmpty() {
		cur, _ := Stack94.Pop()
		res = append(res, cur.Val)
		cur = cur.Right
		for cur != nil {
			Stack94.Push(cur)
			cur = cur.Left
		}
	}
	return res
}

// @lc code=end
