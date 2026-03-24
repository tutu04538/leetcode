/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

package main

import "slices"

// @lc code=start
// func combinationSum(candidates []int, target int) [][]int {
// 	slices.Sort(candidates)
// 	n := len(candidates)

// 	res := [][]int{}
// 	maxLen := target / candidates[0]
// 	curRes := make([]int, maxLen+1)

// 	var dfs39 func(idx int, sum int, curLen int)
// 	dfs39 = func(idx, sum, curLen int) {
// 		if idx == n {
// 			return
// 		}
// 		var next int
// 		if idx < n-1 {
// 			next = candidates[idx+1]
// 		}
// 		for i := range maxLen + 1 {
// 			if i > 0 {
// 				curRes[curLen] = candidates[idx]
// 				sum += candidates[idx]
// 				curLen += 1
// 			}

// 			if sum == target {
// 				tmp := make([]int, curLen)
// 				copy(tmp, curRes)
// 				res = append(res, tmp)
// 				return
// 			}

// 			if sum > target {
// 				return
// 			}

// 			if sum+next > target {
// 				continue
// 			}

// 			dfs39(idx+1, sum, curLen)
// 		}

// 		return
// 	}

// 	dfs39(0, 0, 0)
// 	return res
// }

func combinationSum(candidates []int, target int) [][]int {
	// 1. 排序是剪枝的前提
	slices.Sort(candidates)

	var res [][]int
	var path []int

	// 定义 DFS 闭包
	var dfs func(target int, start int)
	dfs = func(target int, start int) {
		// 基准条件：找到目标
		if target == 0 {
			// 注意：必须拷贝一份 path，因为 path 在回溯中会被修改
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			// 2. 核心剪枝：如果当前数字已经大于剩余 target，
			// 那么由于 candidates 已排序，后续数字也一定大于 target，直接跳出循环
			if candidates[i] > target {
				break
			}

			// 3. 做选择
			path = append(path, candidates[i])

			// 4. 递归：传入 i 而不是 i + 1，表示同一个数可以重复选取
			// 同时更新 target 为 target - candidates[i]
			dfs(target-candidates[i], i)

			// 5. 回溯：撤销选择
			path = path[:len(path)-1]
		}
	}

	dfs(target, 0)
	return res
}

// @lc code=end
