/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

// @lc code=start

// // Stack394 定义了一个泛型栈
// type Stack394[T any] struct {
// 	items []T
// }

// // Push 将元素压入栈顶
// func (s *Stack394[T]) Push(item T) {
// 	s.items = append(s.items, item)
// }

// // Pop 弹出并返回栈顶元素
// func (s *Stack394[T]) Pop() (T, error) {
// 	if len(s.items) == 0 {
// 		var zero T // 返回 T 类型的零值
// 		return zero, errors.New("stack is empty")
// 	}
// 	// 获取最后一个元素
// 	item := s.items[len(s.items)-1]
// 	// 切掉最后一个元素
// 	s.items = s.items[:len(s.items)-1]
// 	return item, nil
// }

// // Peek 查看栈顶元素但不弹出
// func (s *Stack394[T]) Peek() *T {
// 	if len(s.items) == 0 {
// 		return nil
// 	}
// 	return &s.items[len(s.items)-1]
// }

// // IsEmpty 判断栈是否为空
// func (s *Stack394[T]) IsEmpty() bool {
// 	return len(s.items) == 0
// }

// // Size 返回栈的大小
// func (s *Stack394[T]) Size() int {
// 	return len(s.items)
// }

// type repeatStringItem struct {
// 	repeat int
// 	str    *[]byte
// }

// func decodeString(s string) string {

// 	stack394 := Stack394[repeatStringItem]{}

// 	var topStr *[]byte
// 	topStr = &[]byte{}
// 	stack394.Push(repeatStringItem{repeat: 1, str: topStr})

// 	i := 0
// 	for i < len(s) {
// 		if s[i] > '0' && s[i] <= '9' {
// 			repeatNum := 0
// 			for s[i] != '[' {
// 				repeatNum = repeatNum*10 + int(s[i]-'0')
// 				i++
// 			}
// 			topStr = &[]byte{}
// 			stack394.Push(repeatStringItem{repeat: repeatNum, str: topStr})
// 			i++
// 		} else if s[i] == ']' {
// 			// pop and expand
// 			topItem, _ := stack394.Pop()
// 			n := len(*topItem.str)
// 			expandStr := make([]byte, n*topItem.repeat)
// 			for r := range topItem.repeat {
// 				copy(expandStr[r*n:], *topItem.str)
// 			}
// 			// 拼接到当前 top
// 			curTop := stack394.Peek()
// 			result := make([]byte, len(*curTop.str)+len(expandStr))
// 			copy(result, *curTop.str)
// 			copy(result[len(*curTop.str):], expandStr)
// 			curTop.str = &result
// 			// update topStr
// 			topStr = curTop.str
// 			i++
// 		} else {
// 			// ordinary char
// 			*topStr = append(*topStr, s[i])
// 			i++
// 		}
// 	}

// 	return string(*topStr)
// }

func decodeString(s string) string {
	// 两个栈：一个存倍数，一个存当前括号层级之前的字符串
	countStack := []int{}
	resStack := []string{}

	currentStr := ""
	i := 0
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			// 1. 解析数字
			repeatNum := 0
			for s[i] >= '0' && s[i] <= '9' {
				repeatNum = repeatNum*10 + int(s[i]-'0')
				i++
			}
			countStack = append(countStack, repeatNum)
		} else if s[i] == '[' {
			// 2. 进入嵌套：把当前积累的字符串和倍数入栈，开始处理括号内的
			resStack = append(resStack, currentStr)
			currentStr = "" // 重置，准备接收括号内的字符
			i++
		} else if s[i] == ']' {
			// 3. 结束嵌套：弹出倍数和之前的字符串
			repeatNum := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			lastStr := resStack[len(resStack)-1]
			resStack = resStack[:len(resStack)-1]

			// 构建括号内的重复字符串并拼接到上一层
			tmp := ""
			for r := 0; r < repeatNum; r++ {
				tmp += currentStr
			}
			currentStr = lastStr + tmp
			i++
		} else {
			// 4. 普通字母
			currentStr += string(s[i])
			i++
		}
	}
	return currentStr
}

// @lc code=end
