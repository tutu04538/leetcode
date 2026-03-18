/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// if head == nil {
	// 	return nil
	// }
	// toSortSlice := []*ListNode{}
	// current := head
	// for current != nil {
	// 	toSortSlice = append(toSortSlice, current)
	// 	current = current.Next
	// }
	// slices.SortFunc(toSortSlice, func(a *ListNode, b *ListNode) int {
	// 	return a.Val - b.Val
	// })
	// n := len(toSortSlice)
	// toSortSlice[n-1].Next = nil
	// for i := range n - 1 {
	// 	toSortSlice[i].Next = toSortSlice[i+1]
	// }
	// return toSortSlice[0]

	if head == nil {
		return nil
	}
	totLength := 0
	current := head
	for current != nil {
		totLength += 1
		current = current.Next
	}

	mergeList := func(list1 *ListNode, list2 *ListNode) *ListNode {
		dummy := &ListNode{}
		current := dummy
		for list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				current.Next = list1
				list1 = list1.Next
			} else {
				current.Next = list2
				list2 = list2.Next
			}
			current = current.Next
		}
		if list1 == nil {
			current.Next = list2
		} else {
			current.Next = list1
		}
		return dummy.Next
	}
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < totLength; subLength <<= 1 {
		prev, current := dummyHead, dummyHead.Next
		for current != nil {
			left := current
			for range subLength - 1 {
				if current.Next == nil {
					break
				}
				current = current.Next
			}

			right := current.Next
			current.Next = nil
			current = right

			for range subLength - 1 {
				if current == nil || current.Next == nil {
					break
				}
				current = current.Next
			}

			var next *ListNode
			if current != nil {
				next = current.Next
				current.Next = nil
				current = next
			}

			prev.Next = mergeList(left, right)

			for range 2 * subLength {
				if prev.Next == nil {
					break
				}
				prev = prev.Next
			}

			// left, right := current, current
			// var nextSub *ListNode
			// for range subLength - 1 {
			// 	right = right.Next
			// 	if right == nil {
			// 		break
			// 	}
			// }
			// if right == nil || right.Next == nil {
			// 	break
			// }
			// leftTail := right
			// right = right.Next
			// leftTail.Next = nil
			// nextSub = right
			// for range subLength - 1 {
			// 	nextSub = nextSub.Next
			// 	if nextSub == nil {
			// 		break
			// 	}
			// }
			// if nextSub != nil {
			// 	rightTail := nextSub
			// 	nextSub = nextSub.Next
			// 	rightTail.Next = nil
			// }
			// newLeft := mergeList(left, right)
			// prev.Next = newLeft
			// if nextSub == nil {
			// 	break
			// }
			// newTail := newLeft
			// for range subLength*2 - 1 {
			// 	newTail = newTail.Next
			// }
			// prev = newTail
			// current = nextSub
			// prev.Next = current
		}

	}

	return dummyHead.Next
}

// @lc code=end
