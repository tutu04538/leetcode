/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除了自身以外数组的乘积
 */
package main

// @lc code=start
func productExceptSelf(nums []int) []int {
	mulSumLeft := make([]int, len(nums))
	mulSumRight := make([]int, len(nums))
	mulSumLeft[0] = nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		mulSumLeft[i] = mulSumLeft[i-1] * nums[i]
	}
	mulSumRight[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		mulSumRight[i] = mulSumRight[i+1] * nums[i]
	}
	res := make([]int, n)
	res[0] = mulSumRight[1]
	res[n-1] = mulSumLeft[n-2]
	for i := 1; i < n-1; i++ {
		res[i] = mulSumLeft[i-1] * mulSumRight[i+1]
	}
	return res
}

// @lc code=end
