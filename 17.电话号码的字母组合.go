/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package main

// @lc code=start
func letterCombinations(digits string) []string {
	dictionary := make([]string, 8)
	dictionary[0] = "abc"
	dictionary[1] = "def"
	dictionary[2] = "ghi"
	dictionary[3] = "jkl"
	dictionary[4] = "mno"
	dictionary[5] = "pqrs"
	dictionary[6] = "tuv"
	dictionary[7] = "wxyz"

	n := len(digits)
	res := []string{}
	var cur []byte
	cur = make([]byte, n)
	var dfs17 func(idx int)
	dfs17 = func(idx int) {
		if idx == n {
			res = append(res, string(cur))
			return
		}
		for i := range len(dictionary[digits[idx]-'2']) {
			cur[idx] = dictionary[digits[idx]-'2'][i]
			dfs17(idx + 1)
		}
	}

	dfs17(0)
	return res

}

// @lc code=end
