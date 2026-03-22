/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */
package main

// @lc code=start
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		idx := c - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &Trie{}
		}
		cur = cur.children[idx]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		idx := c - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		idx := c - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
