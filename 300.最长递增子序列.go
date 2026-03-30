/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package main

// @lc code=start
func lengthOfLIS(nums []int) int {

	// dp := make([]int, len(nums))

	// dp[0] = 1
	// for i := 1; i < len(nums); i++ {
	// 	dp[i] = 1
	// 	for j := 0; j < i; j++ {
	// 		if nums[i] <= nums[j] {
	// 			continue
	// 		}
	// 		dp[i] = max(dp[i], dp[j]+1)
	// 	}
	// }

	// res := 0
	// for i := 0; i < len(nums); i++ {
	// 	res = max(res, dp[i])
	// }

	// return res

	d := make([]int, len(nums)+1)
	d[1] = nums[0]
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > d[maxLen] {
			maxLen += 1
			d[maxLen] = nums[i]
		} else {
			l, r := 1, maxLen
			target := nums[i]
			for l < r {
				mid := (l + r) >> 1
				if d[mid] < target {
					l = mid + 1
				} else {
					r = mid
				}
			}
			d[l] = target
		}
	}

	return maxLen

}

// @lc code=end
