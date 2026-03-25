/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package main

// @lc code=start
func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	res := make([]int, 2)
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if target != nums[l] {
		res[0], res[1] = -1, -1
		return res
	}
	res[0] = l

	l, r = 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	res[1] = l

	return res

}

// @lc code=end
