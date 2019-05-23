/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 1
	last := head
	for {
		if last.Next != nil {
			length++
			last = last.Next
		} else {
			break
		}
	}

	if length == 0 || k == 0 {
		return head
	}

	k = k % length
	if k == 0 {
		return head
	}

	pre := head
	i := 0
	for {
		if i < length-k-1 {
			pre = pre.Next
			i++
		} else {
			break
		}
	}
	temp := pre.Next

	pre.Next = nil
	last.Next = head
	return temp
}

