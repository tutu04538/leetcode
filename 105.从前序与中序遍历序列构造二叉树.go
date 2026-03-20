/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func recursiveBuildTree(preorder *[]int, inorder []int, inorderBase int, m map[int]int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	curRoot := &TreeNode{Val: (*preorder)[0]}
	rootIdx := m[curRoot.Val] - inorderBase
	// var rootIdx int
	// for idx, val := range inorder {
	// 	if val == curRoot.Val {
	// 		rootIdx = idx
	// 		break
	// 	}
	// }
	*preorder = (*preorder)[1:]
	curRoot.Left = recursiveBuildTree(preorder, inorder[:rootIdx], inorderBase, m)
	curRoot.Right = recursiveBuildTree(preorder, inorder[rootIdx+1:], inorderBase+rootIdx+1, m)
	return curRoot
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for idx, val := range inorder {
		m[val] = idx
	}
	return recursiveBuildTree(&preorder, inorder, 0, m)
}

// @lc code=end
