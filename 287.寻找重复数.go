/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */
package main

// @lc code=start
func findDuplicate(nums []int) int {
	// slow, fast := 0, 0
	// for {
	// 	slow = nums[slow]
	// 	fast = nums[nums[fast]]
	// 	if slow == fast {
	// 		break
	// 	}
	// }
	// slow = 0
	// for {
	// 	slow = nums[slow]
	// 	fast = nums[fast]
	// 	if slow == fast {
	// 		return slow
	// 	}
	// }

	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow

}

// @lc code=end
