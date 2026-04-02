/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	resLen := 1
	resLeft := 0

	for left := 0; left+1 < n; left++ {
		right := left + 1
		if s[left] == s[right] {
			dp[left][right] = true
			resLen = 2
			resLeft = left
		}
	}
	for len := 3; len <= n; len++ {
		for left := 0; left+len-1 < n; left++ {
			right := left + len - 1
			if s[left] == s[right] && dp[left+1][right-1] {
				dp[left][right] = true
				resLen = len
				resLeft = left
			}
		}
	}

	return s[resLeft : resLeft+resLen]
}

// @lc code=end
