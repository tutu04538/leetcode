/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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

type queueItem struct {
	Node  *TreeNode
	Depth int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*queueItem{}
	res := [][]int{}
	queue = append(queue, &queueItem{root, 1})
	for len(queue) > 0 {
		cur := queue[0]
		if len(res) < cur.Depth {
			res = append(res, []int{cur.Node.Val})
		} else {
			res[cur.Depth-1] = append(res[cur.Depth-1], cur.Node.Val)
		}
		queue[0] = nil
		queue = queue[1:]
		if cur.Node.Left != nil {
			queue = append(queue, &queueItem{cur.Node.Left, cur.Depth + 1})
		}
		if cur.Node.Right != nil {
			queue = append(queue, &queueItem{cur.Node.Right, cur.Depth + 1})
		}
	}
	return res
}

// @lc code=end
