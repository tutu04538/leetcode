/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package main

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	res := []int{}
	target := [26]int{}
	current := [26]int{}

	for i := range len(p) {
		target[p[i]-'a']++
		current[s[i]-'a']++
	}

	if target == current {
		res = append(res, 0)
	}

	right := len(p)
	for right < len(s) {
		current[s[right]-'a']++
		current[s[right-len(p)]-'a']--
		if target == current {
			res = append(res, right-len(p)+1)
		}
		right++
	}
	return res
}

// @lc code=end
