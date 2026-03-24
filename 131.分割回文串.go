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
	var check func(curStr string) bool

	check = func(curStr string) bool {
		n := len(curStr)
		for i := range n / 2 {
			if curStr[i] != curStr[n-1-i] {
				return false
			}
		}
		return true
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
			if s[end] == s[curIdx] && check(s[curIdx:end+1]) {
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
