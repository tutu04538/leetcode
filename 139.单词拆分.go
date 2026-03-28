/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */
package main

// @lc code=start
// func wordBreak(s string, wordDict []string) bool {
// 	maxLen := len(wordDict) * 21
// 	trie := make([][]int, maxLen)
// 	isEnd := make([]bool, maxLen)
// 	for i := range maxLen {
// 		trie[i] = make([]int, 26)
// 	}
// 	tot := 0
// 	for i := range len(wordDict) {
// 		curWord := wordDict[i]
// 		p := 0
// 		for j := len(curWord) - 1; j >= 0; j-- {
// 			c := curWord[j] - 'a'
// 			if trie[p][c] != 0 {
// 				p = trie[p][c]
// 			} else {
// 				tot++
// 				trie[p][c] = tot
// 				p = tot
// 			}
// 		}
// 		isEnd[p] = true
// 	}

// 	dp := make([]bool, len(s)+1)
// 	dp[0] = true

// 	for i := 1; i <= len(s); i++ {
// 		p := 0
// 		for j := i; j >= 1; j-- {
// 			c := s[j-1] - 'a'
// 			if trie[p][c] != 0 {
// 				p = trie[p][c]
// 			} else {
// 				break
// 			}
// 			if isEnd[p] && dp[j-1] {
// 				dp[i] = true
// 				break
// 			}
// 		}
// 	}

// 	return dp[len(s)]
// }

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	// 1. 使用一维数组模拟 Trie，极大减少 make 调用
	// 这里的 26000 是根据题目范围预估的节点数
	trie := make([][26]int, 1)
	isEnd := make([]bool, 1)

	// 插入 Trie
	for _, word := range wordDict {
		p := 0
		for i := 0; i < len(word); i++ {
			c := word[i] - 'a'
			if trie[p][c] == 0 {
				trie = append(trie, [26]int{})
				isEnd = append(isEnd, false)
				trie[p][c] = len(trie) - 1
			}
			p = trie[p][c]
		}
		isEnd[p] = true
	}

	dp := make([]bool, n+1)
	dp[0] = true

	// 2. 优化 DP 顺序：只有当前缀合法，才开始匹配
	for i := 0; i < n; i++ {
		if !dp[i] {
			continue
		}
		// 从 i 位置开始在 Trie 中向后匹配
		p := 0
		for j := i; j < n; j++ {
			c := s[j] - 'a'
			if trie[p][c] == 0 {
				break
			}
			p = trie[p][c]
			if isEnd[p] {
				dp[j+1] = true
			}
		}
	}

	return dp[n]
}

// @lc code=end
