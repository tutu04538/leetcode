/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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

func recursiveDiameterOfBinaryTree(cur *TreeNode, res *int) int {
	if cur == nil {
		return 0
	}
	leftMax := recursiveDiameterOfBinaryTree(cur.Left, res)
	rightMax := recursiveDiameterOfBinaryTree(cur.Right, res)
	*res = max(*res, leftMax+rightMax+1)
	return max(leftMax+1, rightMax+1)
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	recursiveDiameterOfBinaryTree(root, &res)
	return res - 1
}

// @lc code=end
