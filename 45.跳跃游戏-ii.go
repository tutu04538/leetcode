/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */
package main

// @lc code=start
func jump(nums []int) int {
	// n := len(nums)
	// if n == 1 {
	// 	return 0
	// }

	// maxFar := make([][2]int, n-1)
	// for i := range n - 1 {
	// 	maxFar[i] = [2]int{i, i + nums[i]}
	// }

	// slices.SortFunc(maxFar, func(a, b [2]int) int {
	// 	if a[1] > b[1] {
	// 		return -1
	// 	} else if a[1] < b[1] {
	// 		return 1
	// 	} else {
	// 		if a[0] < b[0] {
	// 			return -1
	// 		} else {
	// 			return 1
	// 		}
	// 	}
	// })
	// curPostion := n - 1
	// p := 0
	// res := 0
	// for curPostion > 0 {
	// 	nextPostion := curPostion
	// 	for p < n-1 {
	// 		if maxFar[p][0] >= curPostion {
	// 			p++
	// 			continue
	// 		}
	// 		if maxFar[p][1] >= curPostion {
	// 			nextPostion = min(nextPostion, maxFar[p][0])
	// 			p++
	// 			continue
	// 		}
	// 		break
	// 	}
	// 	res += 1
	// 	curPostion = nextPostion
	// }

	// return res
	n := len(nums)
	if n <= 1 {
		return 0
	}

	res := 0      // 跳跃步数
	maxReach := 0 // 目前能跳到的最远位置
	end := 0      // 当前步跳跃的范围边界

	// 注意：我们不需要遍历最后一个元素，
	// 因为题目保证可以到达，如果我们还没到最后就更新了 end，res 就会正确增加。
	for i := 0; i < n-1; i++ {
		// 更新从当前位置及以前能达到的最远点
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}

		// 当我们走到了当前这一步的“极限边界”
		if i == end {
			res++          // 必须再跳一次
			end = maxReach // 下一次跳跃的最远边界更新为 maxReach

			// 如果已经能覆盖到终点，可以直接退出（可选优化）
			if end >= n-1 {
				break
			}
		}
	}

	return res
}

// @lc code=end
