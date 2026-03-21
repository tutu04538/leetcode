/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main

// @lc code=start
type Coordinate struct {
	x, y int
}

func numIslands(grid [][]byte) int {
	res := 0
	// queue := []*Coordinate{}
	m, n := len(grid), len(grid[0])
	// visited := make([][]bool, m)
	// direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '2'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := range m {
		for j := range n {
			if grid[i][j] != '1' {
				continue
			}
			// cur := &Coordinate{i, j}
			// if _, ok := visited[*cur]; ok {
			// 	continue
			// }
			res += 1
			dfs(i, j)
			// visited[*cur] = struct{}{}
			// queue = append(queue, cur)
			// for len(queue) > 0 {
			// 	now := queue[0]
			// 	queue[0] = nil
			// 	queue = queue[1:]
			// 	for _, d := range direction {
			// 		x, y := now.x+d[0], now.y+d[1]
			// 		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '1' {
			// 			continue
			// 		}
			// 		// if _, ok := visited[Coordinate{x, y}]; ok {
			// 		// 	continue
			// 		// }
			// 		queue = append(queue, &Coordinate{x, y})
			// 		// visited[Coordinate{x, y}] = struct{}{}
			// 		grid[x][y] = '2'
			// 	}
			// }
		}
	}
	return res
}

// @lc code=end
