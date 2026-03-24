/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
package main

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	cur := make([]byte, n*2)

	var dfs22 func(idx int, left int, right int)
	dfs22 = func(idx, left, right int) {
		if idx == 2*n {
			res = append(res, string(cur))
			return
		}
		if right > left {
			return
		}
		if left < n {
			cur[idx] = '('
			dfs22(idx+1, left+1, right)
		}
		cur[idx] = ')'
		dfs22(idx+1, left, right+1)
	}

	dfs22(0, 0, 0)
	return res

}

// @lc code=end
