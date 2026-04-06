/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */
package main

// @lc code=start
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, v := range nums {
		if v == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		} else if v == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if nums[i] == 1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		}
	}
}

// @lc code=end
