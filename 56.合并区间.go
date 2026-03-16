/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

package main

import "slices"

// @lc code=start
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})
	res := [][]int{}

	current := slices.Clone(intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > current[1] {
			res = append(res, slices.Clone(current))
			current[0] = intervals[i][0]
			current[1] = intervals[i][1]
			continue
		}
		current[1] = max(current[1], intervals[i][1])
	}

	res = append(res, current)
	return res
}

// @lc code=end
