/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
package main

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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	recursiveInorderTraversal(root, &res)
	return res
}

// @lc code=end
