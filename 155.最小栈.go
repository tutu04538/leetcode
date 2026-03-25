/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */
package main

// @lc code=start
// type MinStack struct {
// 	items   []int
// 	minVals []int
// }

// func Constructor() MinStack {
// 	m := MinStack{}
// 	return m
// }

// func (this *MinStack) Push(val int) {
// 	this.items = append(this.items, val)
// 	if len(this.minVals) == 0 {
// 		this.minVals = append(this.minVals, val)
// 	} else {
// 		if val <= this.minVals[len(this.minVals)-1] {
// 			this.minVals = append(this.minVals, val)
// 		}
// 	}
// }

// func (this *MinStack) Pop() {
// 	item := this.items[len(this.items)-1]
// 	this.items = this.items[:len(this.items)-1]
// 	if item == this.minVals[len(this.minVals)-1] {
// 		this.minVals = this.minVals[:len(this.minVals)-1]
// 	}
// }

// func (this *MinStack) Top() int {
// 	return this.items[len(this.items)-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.minVals[len(this.minVals)-1]
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type node struct {
	val int
	min int
}

type MinStack struct {
	stack []node
}

func Constructor() MinStack {
	return MinStack{
		// 预分配一定的空间可以显著减少 append 带来的内存重新分配
		stack: make([]node, 0, 1000),
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.stack) > 0 {
		currMin := this.stack[len(this.stack)-1].min
		if currMin < min {
			min = currMin
		}
	}
	this.stack = append(this.stack, node{val: val, min: min})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

// @lc code=end
