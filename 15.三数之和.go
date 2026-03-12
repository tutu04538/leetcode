/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

package main

import "slices"

// @lc code=start
func threeSum(nums []int) [][]int {

	/*
		// O(n^2) without sort but with map which is slow
		var i, j int
		m := make(map[int]struct{})
		res := make(map[[3]int]struct{})
		var ans [][]int

		for i = 0; i < len(nums); i++ {
			for j = len(nums) - 1; j > i; j-- {
				thirdNum := -(nums[i] + nums[j])
				if _, ok := m[thirdNum]; ok {
					candidateAns := [3]int{nums[i], nums[j], thirdNum}
					slices.Sort(candidateAns[:])
					res[candidateAns] = struct{}{}
				}
				m[nums[j]] = struct{}{}
			}
			clear(m)
		}

		for k := range res {
			ans = append(ans, k[:])
		}

		return ans
	*/

	// O(n^2) with sort
	slices.Sort(nums)
	var ans [][]int

	i := 0
	for i+2 < len(nums) {
		j, k := i+1, len(nums)-1
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		for j < k {
			target := nums[i] + nums[j] + nums[k]
			if nums[i]+nums[k]+nums[k-1] < 0 {
				break
			}
			if target == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				t1 := nums[j]
				t2 := nums[k]
				for j < k && t1 == nums[j] {
					j++
				}
				for j < k && t2 == nums[k] {
					k--
				}
			} else if target < 0 {
				j++
			} else {
				k--
			}
		}
		t3 := nums[i]
		for i+2 < len(nums) && t3 == nums[i] {
			i++
		}
	}

	return ans

}

// @lc code=end
