/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
package main

// @lc code=start
func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}

	var reverse func(arr []int)
	reverse = func(arr []int) {
		l, r := 0, len(arr)-1
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			continue
		}
		for j := n - 1; j > i; j-- {
			if nums[j] <= nums[i] {
				continue
			}
			nums[j], nums[i] = nums[i], nums[j]
			reverse(nums[i+1:])
			return
		}
	}

	reverse(nums)
	return
}

// @lc code=end
