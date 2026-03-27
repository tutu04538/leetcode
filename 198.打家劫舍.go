/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package main

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	if n == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	if n == 2 {
		return dp[1]
	}

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	return dp[n-1]
}

// @lc code=end
