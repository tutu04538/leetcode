/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

package main

// @lc code=start
func subarraySum(nums []int, k int) int {
	res := 0
	sum := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			sum[i] = v
		} else {
			sum[i] = sum[i-1] + v
		}
	}

	m := make(map[int]int)
	if nums[len(nums)-1] == k {
		res += 1
	}
	m[sum[len(nums)-1]] += 1
	for i := len(nums) - 2; i >= 0; i-- {
		// nums[i] + sum[x] - sum[i] == k
		m[sum[i]] += 1
		target := k + sum[i] - nums[i]
		res += m[target]
	}

	return res
}

// @lc code=end
