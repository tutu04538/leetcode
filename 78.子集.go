/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

package main

// @lc code=start
func subsets(nums []int) [][]int {

	n := len(nums)
	res := [][]int{}
	curRes := make([]int, n)
	var dfs78 func(idx, len int)
	dfs78 = func(idx, curLen int) {
		if idx == n {
			tmp := make([]int, curLen)
			copy(tmp, curRes)
			res = append(res, tmp)
			return
		}
		dfs78(idx+1, curLen)
		curRes[curLen] = nums[idx]
		dfs78(idx+1, curLen+1)
	}

	dfs78(0, 0)

	return res

}

// @lc code=end
