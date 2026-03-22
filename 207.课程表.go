/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */
package main

// @lc code=start

type Edge struct {
	to, weight, next int
}

var cnt int
var edges []Edge
var heads []int

func add_edge(u, v int) {
	edges[cnt] = Edge{to: v, next: heads[u]}
	heads[u] = cnt
	cnt++
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	cnt = 1
	edges = make([]Edge, len(prerequisites)+1)
	heads = make([]int, numCourses)

	for _, item := range prerequisites {
		add_edge(item[0], item[1])
	}

	states := make([]int, numCourses)

	var dfs207 func(cur int) bool
	dfs207 = func(cur int) bool {
		states[cur] = 1
		for i := heads[cur]; i != 0; i = edges[i].next {
			v := edges[i].to
			if states[v] == 1 {
				return false
			}
			if states[v] == 0 {
				if !dfs207(v) {
					return false
				}
			}
		}
		states[cur] = 2
		return true
	}

	for i := range numCourses {
		if states[i] == 0 {
			if !dfs207(i) {
				return false
			}
		}
	}

	return true

}

// @lc code=end
