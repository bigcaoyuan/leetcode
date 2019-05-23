/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		root *ListNode
		node *ListNode
		temp *ListNode
	)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		root = l1
		l1 = l1.Next
	} else {
		root = l2
		l2 = l2.Next
	}
	node = root
	for {
		if l1 == nil {
			node.Next = l2
			break
		}
		if l2 == nil {
			node.Next = l1
			break
		}

		if l1.Val <= l2.Val {
			temp = l1.Next
			node.Next = l1
			l1 = temp
			node = node.Next
		} else {
			temp = l2.Next
			node.Next = l2
			l2 = temp
			node = node.Next
		}
	}
	return root
}

