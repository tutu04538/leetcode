/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

import "math/rand"

// @lc code=start

func findKthLargest(nums []int, k int) int {

	// rand.Seed(time.Now().UnixNano())

	// var quickSort func(arr []int, _k int) int
	// var partition215 func(arr []int) int

	// quickSort = func(arr []int, _k int) int {
	// 	index := partition215(arr)
	// 	curBiggerTot := index + 1
	// 	switch {
	// 	case curBiggerTot == _k:
	// 		return arr[index]
	// 	case curBiggerTot > _k:
	// 		return quickSort(arr[:index], _k)
	// 	default:
	// 		return quickSort(arr[index+1:], _k-curBiggerTot)
	// 	}
	// }

	// partition215 = func(arr []int) int {
	// 	pivot := arr[0]
	// 	index := 1
	// 	for i := 1; i < len(arr); i++ {
	// 		if arr[i] > pivot {
	// 			arr[i], arr[index] = arr[index], arr[i]
	// 			index++
	// 		}
	// 	}
	// 	index--
	// 	arr[index], arr[0] = arr[0], arr[index]
	// 	return index
	// }

	// return quickSort(nums, k)

	var partition215 func(left, right int) int

	partition215 = func(left, right int) int {
		randomIndex := rand.Intn(right-left+1) + left
		pivot := nums[randomIndex]
		nums[randomIndex], nums[left] = nums[left], nums[randomIndex]
		i, j := left+1, right
		for {
			for i <= j && nums[i] > pivot {
				i++
			}

			for i <= j && nums[j] < pivot {
				j--
			}

			if i >= j {
				break
			}

			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}

		nums[left], nums[j] = nums[j], nums[left]
		return j
	}

	target := k - 1
	left, right := 0, len(nums)-1
	for left <= right {
		index := partition215(left, right)
		if index == target {
			return nums[index]
		} else if index > target {
			right = index - 1
		} else {
			left = index + 1
		}
	}

	return -1

}

// @lc code=end
