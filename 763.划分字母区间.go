/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */
package main

// @lc code=start
func partitionLabels(s string) []int {
	maxChar := [26]int{}
	for i := 0; i < len(s); i++ {
		maxChar[s[i]-'a'] = i
	}
	res := []int{}
	curEnd := -1
	lastEnd := -1
	cur := 0
	n := len(s)
	for cur < n {
		curEnd = max(curEnd, maxChar[s[cur]-'a'])
		if cur == curEnd {
			res = append(res, curEnd-lastEnd)
			lastEnd = curEnd
		}
		cur++
	}
	return res
}

// @lc code=end
