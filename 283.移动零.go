/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

// @lc code=start
func moveZeroes(nums []int) {
	p1, p2 := 0, 0
	for _, v := range nums {
		if v != 0 {
			nums[p1] = nums[p2]
			p1++
			p2++
		} else {
			p2++
		}
	}

	for ; p1 < len(nums); p1++ {
		nums[p1] = 0
	}

}

// @lc code=end
