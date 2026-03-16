/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */
package main

// @lc code=start

type state struct {
	lSum, rSum, allSum, maxSum int
}

func calSubArray(l int, r int, nums []int) state {
	if l == r {
		return state{nums[l], nums[l], nums[l], nums[l]}
	}
	mid := (l + r) / 2
	stateL := calSubArray(l, mid, nums)
	stateR := calSubArray(mid+1, r, nums)
	stateC := state{}
	stateC.allSum = stateL.allSum + stateR.allSum
	stateC.lSum = max(stateL.lSum, stateL.allSum+stateR.lSum)
	stateC.rSum = max(stateR.rSum, stateR.allSum+stateL.rSum)
	stateC.maxSum = max(stateL.maxSum, stateR.maxSum, stateL.rSum+stateR.lSum)
	return stateC
}

func maxSubArray(nums []int) int {
	/*
		// simple O(n)
		sum := make([]int, len(nums))
		sum[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			sum[i] = sum[i-1] + nums[i]
		}

		min_sum := 0
		res := math.MinInt32
		for i := 0; i < len(nums); i++ {
			res = max(res, sum[i]-min_sum)
			min_sum = min(sum[i], min_sum)
		}

		return res
	*/

	// 分治（线段树）
	stateA := calSubArray(0, len(nums)-1, nums)
	return stateA.maxSum

}

// @lc code=end
