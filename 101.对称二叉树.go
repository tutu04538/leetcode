/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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

func recursiveIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	if !recursiveIsSymmetric(left.Left, right.Right) || !recursiveIsSymmetric(left.Right, right.Left) {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	return recursiveIsSymmetric(root.Left, root.Right)
}

// @lc code=end
