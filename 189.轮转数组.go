/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

package main

// @lc code=start
func rotate(nums []int, k int) {
	/*
		// 空间复杂度 O(N)
		k %= len(nums)
		if k == 0 {
			return
		}
		ans := nums[(len(nums) - k):]
		for i := 0; i < len(nums)-k; i++ {
			ans = append(ans, nums[i])
		}
		copy(nums, ans)
	*/
	k %= len(nums)
	if k == 0 {
		return
	}

	reverse := func(s []int) {
		for i := range len(s) / 2 {
			s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
		}
	}

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

	// for i := range len(nums) / 2 {
	// 	nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	// }
	// for i := range k / 2 {
	// 	nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	// }

	// for i := k; i < (len(nums)-1+k)/2; i++ {
	// 	nums[i], nums[k+len(nums)-1-i] = nums[k+len(nums)-1-i], nums[i]
	// }

}

// @lc code=end
