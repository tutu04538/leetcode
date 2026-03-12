/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

package main

// @lc code=start
func maxArea(height []int) int {
	// 双指针，本质上是简化两个 for 循环，类似剪枝
	// 假设外层循环是 x，从小到大，内层循环是 y，从大到小，那么当 h[x] < h[y] 时，y 再减小就没有意义，直接跳过，让 x 加 1
	// 分情况讨论就可以发现 y 也没必要从末尾开始遍历
	left, right, output := 0, len(height)-1, 0
	for left < right {
		if height[left] <= height[right] {
			output = max(output, height[left]*(right-left))
			left++
		} else {
			output = max(output, height[right]*(right-left))
			right--
		}
	}
	return output
}

// @lc code=end
