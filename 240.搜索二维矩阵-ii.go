/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for {
		if i >= m || j < 0 {
			return false
		}
		if target == matrix[i][j] {
			return true
		}
		if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}
}

// @lc code=end
