/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package main

import "errors"

// @lc code=start

// Stack20 定义了一个泛型栈
type Stack20[T any] struct {
	items []T
}

// Push 将元素压入栈顶
func (s *Stack20[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop 弹出并返回栈顶元素
func (s *Stack20[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T // 返回 T 类型的零值
		return zero, errors.New("stack is empty")
	}
	// 获取最后一个元素
	item := s.items[len(s.items)-1]
	// 切掉最后一个元素
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek 查看栈顶元素但不弹出
func (s *Stack20[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty 判断栈是否为空
func (s *Stack20[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 返回栈的大小
func (s *Stack20[T]) Size() int {
	return len(s.items)
}

func isValid(s string) bool {
	stack := Stack20[byte]{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
			continue
		}

		if c == ')' {
			top, err := stack.Pop()
			if err != nil {
				return false
			}
			if top != '(' {
				return false
			}
		}

		if c == ']' {
			top, err := stack.Pop()
			if err != nil {
				return false
			}
			if top != '[' {
				return false
			}
		}

		if c == '}' {
			top, err := stack.Pop()
			if err != nil {
				return false
			}
			if top != '{' {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

// @lc code=end
