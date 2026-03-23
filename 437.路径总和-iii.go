/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func dfs437(cur *TreeNode, currentSum int, m map[int64]int, res *int, targetSum int) {
	if cur == nil {
		return
	}
	// sum[current] - sum[x] == targetSum
	// sum[x] == sum[cuurent] - targetSum
	currentSum += cur.Val
	target := int64(currentSum) - int64(targetSum)
	*res += m[target]
	m[int64(currentSum)]++
	dfs437(cur.Left, currentSum, m, res, targetSum)
	dfs437(cur.Right, currentSum, m, res, targetSum)
	m[int64(currentSum)]--
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	m := make(map[int64]int)
	m[0] = 1
	dfs437(root, 0, m, &res, targetSum)
	return res
}

// @lc code=end
