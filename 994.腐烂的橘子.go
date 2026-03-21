/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */
package main

// @lc code=start
// func orangesRotting(grid [][]int) int {

// 	m, n := len(grid), len(grid[0])
// 	queue := []int{}
// 	visited := make([][]int, m)
// 	for i := range m {
// 		visited[i] = make([]int, n)
// 	}

// 	normal := 0
// 	abnormal := 0

// 	for i := range m {
// 		for j := range n {
// 			if grid[i][j] == 2 {
// 				queue = append(queue, i*n+j)
// 				visited[i][j] = 1
// 				abnormal++
// 			}
// 			if grid[i][j] == 1 {
// 				normal++
// 			}
// 		}
// 	}

// 	if normal == 0 {
// 		return 0
// 	}

// 	if len(queue) == 0 {
// 		return -1
// 	}

// 	res := 0
// 	last := 1
// 	direction := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
// 	for len(queue) > 0 {
// 		cur := queue[0]
// 		queue = queue[1:]
// 		curX, curY := cur/n, cur%n
// 		if visited[curX][curY] != last {
// 			res++
// 			last = visited[curX][curY]
// 		}
// 		for _, d := range direction {
// 			nextX, nextY := curX+d[0], curY+d[1]
// 			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || grid[nextX][nextY] == 0 || visited[nextX][nextY] != 0 {
// 				continue
// 			}
// 			visited[nextX][nextY] = visited[curX][curY] + 1
// 			queue = append(queue, nextX*n+nextY)
// 		}
// 	}
// 	for i := range m {
// 		for j := range n {
// 			if grid[i][j] == 1 && visited[i][j] == 0 {
// 				return -1
// 			}
// 		}
// 	}
// 	return res
// }

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fresh := 0
	queue := make([]int, 0, m*n) // 预分配空间

	// 1. 统计新鲜橘子，并将初始腐烂橘子入队
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, i*n+j)
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	// 如果没有新鲜橘子，直接返回 0
	if fresh == 0 {
		return 0
	}

	minutes := 0
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// 2. 层级 BFS
	for len(queue) > 0 {
		// 如果 fresh 已经为 0，说明已经全部腐烂，不需要再计时下一分钟
		if fresh == 0 {
			return minutes
		}
		minutes++

		// 处理当前分钟内“新产生”的所有腐烂橘子
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			r, c := curr/n, curr%n

			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				// 越界或不是新鲜橘子则跳过
				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
					grid[nr][nc] = 2 // 直接修改原数组，标记为腐烂
					fresh--          // 新鲜橘子减少
					queue = append(queue, nr*n+nc)
				}
			}
		}
	}

	// 3. 最终判断
	if fresh > 0 {
		return -1
	}
	return minutes
}

// @lc code=end
