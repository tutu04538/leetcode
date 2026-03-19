/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第 K 小的元素
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

var ans int

func traverse230(curNode *TreeNode, curK *int) {
	if curNode == nil {
		return
	}
	traverse230(curNode.Left, curK)
	if *curK < 0 {
		return
	}
	*curK--
	if *curK == 0 {
		ans = curNode.Val
		*curK = -1
		return
	}
	traverse230(curNode.Right, curK)
}

func kthSmallest(root *TreeNode, k int) int {
	cur := k
	traverse230(root, &cur)
	return ans
}

// @lc code=end
