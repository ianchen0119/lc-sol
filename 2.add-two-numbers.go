/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := l1
	prev := res
	newVal := 0
	isFirst := true
	c := 0
	for l1 != nil || l2 != nil || c != 0 {
		if l1 != nil && l2 != nil {
			newVal = l1.Val + l2.Val + c
			if newVal >= 10 {
				newVal -= 10
				c = 1
			} else {
				c = 0
			}
			l1.Val = newVal
			prev = l1
		} else if l1 != nil {
			// l2 is nil
			newVal = l1.Val + c
			if newVal >= 10 {
				newVal -= 10
				c = 1
			} else {
				c = 0
			}
			l1.Val = newVal
			prev = l1
		} else if l2 != nil {
			// l1 is nil
			if isFirst {
				prev.Next = l2
				isFirst = false
			}
			newVal = l2.Val + c
			if newVal >= 10 {
				newVal -= 10
				c = 1
			} else {
				c = 0
			}
			l2.Val = newVal
			prev = l2
		} else {
			prev.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
			return res
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return res
}

// @lc code=end

