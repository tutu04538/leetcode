/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)

	binarySearch := func(toSearchNums []int, toSearchTarget int) int {
		l, r := 0, len(toSearchNums)-1
		for l < r {
			mid := (l + r) / 2
			if toSearchTarget > toSearchNums[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if toSearchTarget == toSearchNums[l] {
			return l
		} else {
			return -1
		}
	}

	if nums[0] <= nums[n-1] {
		return binarySearch(nums, target)
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

	if target >= nums[0] {
		return binarySearch(nums[:l], target)
	} else {
		res := binarySearch(nums[l:], target)
		if res == -1 {
			return -1
		} else {
			return l + res
		}
	}
}

// @lc code=end
