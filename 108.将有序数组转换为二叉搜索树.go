/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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

func recusiveBuildBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	if len(nums) == 0 {
		return nil
	}
	root_idx := len(nums) / 2
	leftNode := recusiveBuildBST(nums[:root_idx])
	rightNode := recusiveBuildBST(nums[root_idx+1:])
	root := &TreeNode{nums[root_idx], leftNode, rightNode}
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return recusiveBuildBST(nums)
}

// @lc code=end
