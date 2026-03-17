/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package main

// @lc code=start
func rotate(matrix [][]int) {
	// 向右旋转 90 度 == 先转置再水平翻转
	m, n := len(matrix), len(matrix[0])
	for i := range m {
		for j := range i {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := range m {
		for j := range n / 2 {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

// @lc code=end
