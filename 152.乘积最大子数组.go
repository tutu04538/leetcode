/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

package main

// @lc code=start
func maxProduct(nums []int) int {
	CMin, CMax, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		c := nums[i]
		tmpCMin := CMin
		CMin = min(CMin*c, CMax*c, c)
		CMax = max(CMax*c, tmpCMin*c, c)
		res = max(res, CMax)
	}

	return res

}

// @lc code=end
