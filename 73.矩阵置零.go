/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

package main

// @lc code=start
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make(map[int]struct{}), make(map[int]struct{})
	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}

	for i := range row {
		for j := range n {
			matrix[i][j] = 0
		}
	}
	for i := range col {
		for j := range m {
			matrix[j][i] = 0
		}
	}
}

// @lc code=end
