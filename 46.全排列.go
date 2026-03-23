/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package main

// @lc code=start
func permute(nums []int) [][]int {

	res := [][]int{}
	visited := make([]bool, len(nums))
	n := len(nums)
	cur_res := make([]int, n)

	var dfs46 func(cur_idx int)
	dfs46 = func(cur_idx int) {
		if cur_idx == n {
			tmp := make([]int, n)
			copy(tmp, cur_res)
			res = append(res, tmp)
			return
		}
		for i := range n {
			if visited[i] {
				continue
			}
			visited[i] = true
			cur_res[cur_idx] = nums[i]
			dfs46(cur_idx + 1)
			visited[i] = false
		}
	}

	dfs46(0)
	return res
}

// @lc code=end
