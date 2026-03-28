/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

package main

import "slices"

// @lc code=start
func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	slices.Sort(coins)

	// n := len(coins)
	// for i := 1; i <= amount; i++ {
	// 	dp[i] = 1e9
	// 	for j := 0; j < n && coins[j] <= i; j++ {
	// 		dp[i] = min(dp[i], dp[i-coins[j]]+1)
	// 	}
	// }

	for i := range dp {
		dp[i] = 1e9
	}
	dp[0] = 0
	for _, x := range coins {
		for c := x; c <= amount; c++ {
			dp[c] = min(dp[c], dp[c-x]+1)
		}
	}

	if dp[amount] == 1e9 {
		return -1
	}
	return dp[amount]

}

// @lc code=end
