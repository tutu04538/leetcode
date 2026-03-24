/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package main

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	wordLen := len(word)
	direction := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := make([]bool, m*n)

	var dfs79 func(toCheckCharIdx int, coordinate int) bool

	dfs79 = func(toCheckCharIdx, coordinate int) bool {
		x, y := coordinate/n, coordinate%n
		if board[x][y] != word[toCheckCharIdx] {
			return false
		}
		if toCheckCharIdx == wordLen-1 {
			return true
		}

		visited[coordinate] = true
		for i := range direction {
			nx, ny := x+direction[i][0], y+direction[i][1]
			nextCoordinate := nx*n + ny
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nextCoordinate] {
				if dfs79(toCheckCharIdx+1, nextCoordinate) {
					return true
				}
			}
		}
		visited[coordinate] = false
		return false
	}

	for i := range m {
		for j := range n {
			if board[i][j] == word[0] {
				if dfs79(0, i*n+j) {
					return true
				}
			}
		}
	}

	return false
}

// @lc code=end
