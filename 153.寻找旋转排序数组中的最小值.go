/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package main

// @lc code=start
func findMin(nums []int) int {
	n := len(nums)

	if nums[0] <= nums[n-1] {
		return nums[0]
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= nums[0] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

// @lc code=end
