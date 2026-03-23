/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func dfs236_1(cur *TreeNode, father map[*TreeNode]*TreeNode) {
	if cur == nil {
		return
	}
	if cur.Left != nil {
		father[cur.Left] = cur
	}
	if cur.Right != nil {
		father[cur.Right] = cur
	}
	dfs236_1(cur.Left, father)
	dfs236_1(cur.Right, father)
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

// 	father := make(map[*TreeNode]*TreeNode)
// 	flag := make(map[*TreeNode]struct{})
// 	father[root] = root
// 	dfs236_1(root, father)
// 	for {
// 		flag[p] = struct{}{}
// 		if p == father[p] {
// 			break
// 		}
// 		p = father[p]
// 	}
// 	for {
// 		if _, ok := flag[q]; ok {
// 			return q
// 		}
// 		q = father[q]
// 	}
// }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}

// @lc code=end
