/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func recursiveTraverse(curNode *TreeNode, curLen int, res *int) {
	if curNode == nil {
		*res = max(*res, curLen)
		return
	}
	recursiveTraverse(curNode.Left, curLen+1, res)
	recursiveTraverse(curNode.Right, curLen+1, res)
}

func maxDepth(root *TreeNode) int {
	res := 0
	recursiveTraverse(root, 0, &res)
	return res
}

// @lc code=end
