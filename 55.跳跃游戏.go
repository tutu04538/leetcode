/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package main

// @lc code=start
func canJump(nums []int) bool {
	curMax := -1
	for i := 0; i < len(nums)-1; i++ {
		curMax = max(curMax, i+nums[i])
		if nums[i] == 0 && curMax <= i {
			return false
		}
		if curMax >= len(nums)-1 {
			return true
		}
	}
	return true
}

// @lc code=end
