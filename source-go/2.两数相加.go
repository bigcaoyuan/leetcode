/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	pre := root
	add := 0
	l1V := 0
	l2V := 0
	for {
		if l1 == nil {
			l1V = 0
			if l2 == nil {
				break
			}
		} else {
			l1V = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			l2V = 0
		} else {
			l2V = l2.Val
			l2 = l2.Next
		}
		val := l1V + l2V + add
		if val >= 10 {
			val = val - 10
			add = 1
		} else {
			add = 0
		}
		tempNode := &ListNode{Val: val}
		pre.Next = tempNode
		pre = tempNode
	}
	if add == 1 {
		pre.Next = &ListNode{Val: 1}
	}
	return root.Next
}