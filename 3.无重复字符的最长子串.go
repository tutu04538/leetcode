/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}
	/*
		// O(N) but slow
		var left, right int = 0, 1
		var res int = 1
		m := make(map[byte]struct{})
		m[s[left]] = struct{}{}
		for right < len(s) {
			if _, ok := m[s[right]]; ok {
				for left < right {
					if s[left] == s[right] {
						left++
						break
					}
					delete(m, s[left])
					left++
				}
			} else {
				res = max(res, right-left+1)
				m[s[right]] = struct{}{}
			}
			right++
		}
		return res
	*/

	// fast solution
	// 直接记录每个字符上一次出现的位置，这样 left 可以直接跳转，不用再一次次累加
	var idx [128]int
	for i := range idx {
		idx[i] = -1
	}

	var left, right int = 0, 1
	var res int = 1
	idx[s[left]] = 0
	for right < len(s) {
		if idx[s[right]] >= left {
			left = idx[s[right]] + 1
		} else {
			res = max(res, right-left+1)
		}
		idx[s[right]] = right
		right++
	}

	return res

}

// @lc code=end
