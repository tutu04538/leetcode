/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */
package main

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	res := make([]int, len(temperatures))

	top := -1
	for i, t := range temperatures {
		for top >= 0 && t > temperatures[stack[top]] {
			res[stack[top]] = i - stack[top]
			top--
		}
		top++
		stack[top] = i
	}

	for top >= 0 {
		res[stack[top]] = 0
		top--
	}

	return res
}

// @lc code=end
