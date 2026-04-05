/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package main

import "slices"

// @lc code=start
func majorityElement(nums []int) int {
	slices.Sort(nums)
	return nums[len(nums)/2]
}

// @lc code=end
