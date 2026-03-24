/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */
package main

// @lc code=start
func partition(s string) [][]string {

	res := [][]string{}
	strLen := len(s)

	var dfs131 func(curIdx int, curRes []string)

	dp := make([][]bool, strLen)
	for i := range strLen {
		dp[i] = make([]bool, strLen)
	}
	for i := range strLen {
		dp[i][i] = true
	}

	for len := 2; len <= strLen; len++ {
		for left := 0; left+len-1 < strLen; left++ {
			right := left + len - 1
			if left+1 > right-1 {
				dp[left+1][right-1] = true
			}
			dp[left][right] = dp[left+1][right-1] && (s[left] == s[right])
		}
	}

	dfs131 = func(curIdx int, curRes []string) {
		if curIdx == strLen {
			tmp := make([]string, len(curRes))
			copy(tmp, curRes)
			res = append(res, tmp)
			return
		}
		// only self
		curRes = append(curRes, s[curIdx:curIdx+1])
		dfs131(curIdx+1, curRes)
		curRes = curRes[:len(curRes)-1]
		for end := curIdx + 1; end < strLen; end++ {
			if dp[curIdx][end] {
				curRes = append(curRes, s[curIdx:end+1])
				dfs131(end+1, curRes)
				curRes = curRes[:len(curRes)-1]
			}
		}
		return
	}

	dfs131(0, []string{})

	return res

}

// @lc code=end
