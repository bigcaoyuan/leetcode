/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	nextNext := head.Next.Next
	head.Next = nextNext
	next.Next = head
	pre := head
	head = next

	temp := pre.Next
	for {
		if temp == nil || temp.Next == nil {
			break
		}
		next = temp.Next
		nextNext := temp.Next.Next
		pre.Next = next
		next.Next = temp
		temp.Next = nextNext
		pre = temp
		temp = pre.Next
	}

	return head
}

