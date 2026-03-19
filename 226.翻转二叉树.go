/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */

package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recursiveInvertTree(cur *TreeNode) {
	if cur == nil {
		return
	}
	recursiveInvertTree(cur.Left)
	recursiveInvertTree(cur.Right)
	right := cur.Right
	cur.Right = cur.Left
	cur.Left = right
	return
}

func invertTree(root *TreeNode) *TreeNode {
	recursiveInvertTree(root)
	return root
}

// @lc code=end
