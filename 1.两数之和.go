/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {

	// O(n^2)
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// O(n)
	m1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v1 := nums[i]
		v2 := target - v1
		if idx, ok := m1[v2]; ok {
			return []int{i, idx}
		} else {
			m1[v1] = i
		}
	}

	return []int{}
}

// @lc code=end
