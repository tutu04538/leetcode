/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package main

import "math/rand"

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int, len(nums))
	for _, v := range nums {
		frequencyMap[v]++
	}

	frequencyNum := make([]int, len(frequencyMap))

	i := 0
	for _, v := range frequencyMap {
		frequencyNum[i] = v
		i++
	}

	var _quickSort347 func(arr []int, left, right int) int
	_quickSort347 = func(arr []int, left, right int) int {
		randomIndex := rand.Intn(right-left+1) + left
		pivot := arr[randomIndex]
		arr[randomIndex], arr[left] = arr[left], arr[randomIndex]
		i, j := left+1, right
		for {
			for i <= j && arr[i] >= pivot {
				i++
			}
			for i <= j && arr[j] <= pivot {
				j--
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
		arr[left], arr[j] = arr[j], arr[left]
		return j
	}

	left, right := 0, len(frequencyNum)-1
	target := k - 1
	var minFrequency int
	for left <= right {
		index := _quickSort347(frequencyNum, left, right)
		if index == target {
			minFrequency = frequencyNum[index]
			break
		} else if index > target {
			right = index - 1
		} else {
			left = index + 1
		}
	}

	res := []int{}
	for k, v := range frequencyMap {
		if v >= minFrequency {
			res = append(res, k)
		}
	}

	return res

}

// @lc code=end
