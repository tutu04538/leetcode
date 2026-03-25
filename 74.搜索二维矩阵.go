/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l < r {
		mid := (l + r) / 2
		x, y := mid/n, mid%n
		if target > matrix[x][y] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	x, y := l/n, l%n
	if target == matrix[x][y] {
		return true
	}
	return false
}

// @lc code=end
