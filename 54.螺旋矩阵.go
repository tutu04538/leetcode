/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	flag := make([][]bool, m)
	for i := range m {
		flag[i] = make([]bool, n)
	}
	res := []int{}
	i, j := 0, 0
	k := 0
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for {
		res = append(res, matrix[i][j])
		flag[i][j] = true
		xd, yd := i+direction[k][0], j+direction[k][1]
		if xd >= 0 && xd < m && yd >= 0 && yd < n && !flag[xd][yd] {
			i = xd
			j = yd
			continue
		}
		k = (k + 1) % 4
		xd, yd = i+direction[k][0], j+direction[k][1]
		if xd >= 0 && xd < m && yd >= 0 && yd < n && !flag[xd][yd] {
			i = xd
			j = yd
			continue
		} else {
			break
		}

	}
	return res
}

// @lc code=end
