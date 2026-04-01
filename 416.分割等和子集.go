/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package main

// @lc code=start
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	maxNum := 0
	for _, v := range nums {
		sum += v
		maxNum = max(maxNum, v)
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	if maxNum > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true
	dp[nums[0]] = true

	for i := 1; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]

}

// @lc code=end
