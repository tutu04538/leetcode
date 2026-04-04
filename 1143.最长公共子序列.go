/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */
package main

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {

	m, n := len(text1), len(text2)
	if m < n {
		text1, text2 = text2, text1
		m, n = n, m
	}

	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		prev := 0
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = tmp
		}
	}

	return dp[n]

}

// @lc code=end
