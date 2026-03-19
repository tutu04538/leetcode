/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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

type queueItem199 struct {
	Node  *TreeNode
	Depth int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*queueItem199{}
	res := []int{}

	queue = append(queue, &queueItem199{root, 1})
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if len(res) < cur.Depth {
			res = append(res, cur.Node.Val)
		} else {
			res[cur.Depth-1] = cur.Node.Val
		}
		if cur.Node.Left != nil {
			queue = append(queue, &queueItem199{cur.Node.Left, cur.Depth + 1})
		}
		if cur.Node.Right != nil {
			queue = append(queue, &queueItem199{cur.Node.Right, cur.Depth + 1})
		}
	}
	return res
}

// @lc code=end
