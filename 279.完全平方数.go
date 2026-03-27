/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */
package main

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)

	// nums := []int{}
	// for i := 1; i*i <= n; i++ {
	// 	nums = append(nums, i*i)
	// }

	// var dfs279 func(x int)
	// dfs279 = func(x int) {
	// 	if x == 0 {
	// 		return
	// 	}
	// 	if dp[x] == 0 {
	// 		dp[x] = 1e9
	// 	} else {
	// 		return
	// 	}

	// 	for i := range len(nums) {
	// 		y := x - nums[i]
	// 		if y < 0 {
	// 			break
	// 		}
	// 		dfs279(y)
	// 		dp[x] = min(dp[x], dp[y]+1)
	// 	}
	// }

	// dfs279(n)

	// dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // 最坏情况全是 1*1，初始化为 i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

// @lc code=end
