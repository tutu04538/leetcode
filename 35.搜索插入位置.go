/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
package main

// @lc code=start
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l < r {
		mid = (l + r) / 2
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if target == nums[l] {
		return l
	}
	if target > nums[l] {
		return l + 1
	} else {
		return l
	}
}

// @lc code=end
