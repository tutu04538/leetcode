/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

type state98 struct {
	MinV, MaxV int
	isValidBST bool
}

func recursiveIsVaildBST(cur *TreeNode) *state98 {
	if cur == nil {
		return nil
	}
	var MinV, MaxV int
	leftState := recursiveIsVaildBST(cur.Left)
	if leftState != nil {
		if !leftState.isValidBST || cur.Val <= leftState.MaxV {
			return &state98{
				isValidBST: false,
			}
		} else {
			MinV = leftState.MinV
		}
	} else {
		MinV = cur.Val
	}
	rightState := recursiveIsVaildBST(cur.Right)
	if rightState != nil {
		if !rightState.isValidBST || cur.Val >= rightState.MinV {
			return &state98{
				isValidBST: false,
			}
		} else {
			MaxV = rightState.MaxV
		}
	} else {
		MaxV = cur.Val
	}
	return &state98{
		isValidBST: true,
		MinV:       MinV,
		MaxV:       MaxV,
	}
}

func isValidBST(root *TreeNode) bool {
	return recursiveIsVaildBST(root).isValidBST
}

// @lc code=end
