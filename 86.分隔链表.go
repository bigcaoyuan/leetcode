/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	mid := false
	npre := &ListNode{x - 1, head}
	midpre := npre
	pre := npre
	cur := head
	for {
		if cur == nil {
			break
		}
		if cur.Val < x {
			if !mid {
				midpre = cur
				cur = cur.Next
			} else {
				pre.Next = cur.Next
				cur.Next = midpre.Next
				midpre.Next = cur
				midpre = midpre.Next
				cur = pre.Next
			}
		} else {
			if !mid {
				mid = true
			}
			pre = cur
			cur = cur.Next
		}
	}
	return npre.Next
}

