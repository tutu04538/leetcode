/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

package main

// @lc code=start
func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int, len(nums))

	var res int = 1

	for _, v := range nums {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = 1
		left, ok1 := m[v-1]
		right, ok2 := m[v+1]

		if ok1 && !ok2 {
			m[v-left] = left + 1
			m[v] = left + 1
			res = max(res, left+1)
		}

		if ok2 && !ok1 {
			m[v+right] = right + 1
			m[v] = right + 1
			res = max(res, right+1)
		}

		if ok1 && ok2 {
			m[v-left] = left + right + 1
			m[v+right] = left + right + 1
			res = max(res, left+right+1)
		}
	}
	return res
}

// @lc code=end
