/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package main

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}

	if n == 0 {
		return m
	}

	if m == 0 {
		return n
	}

	for i := 1; i <= m; i++ {
		prev := i - 1
		dp[0] = i
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev
			} else {
				dp[j] = min(dp[j]+1, dp[j-1]+1, prev+1)
			}
			prev = tmp
		}
	}

	return dp[n]
}

// @lc code=end
