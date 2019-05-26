/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	if head == nil {
		return head
	}
	npre := &ListNode{Next: head}
	pre := npre
	for i := 1; i < m; i++ {
		if i < m {
			pre = pre.Next
		}
	}
	temp := pre.Next
	next := temp.Next
	for i := m; i < n; i++ {
		temp.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

		next = temp.Next
	}
	return npre.Next
}

