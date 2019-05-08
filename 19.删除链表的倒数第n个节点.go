/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	pre := head
	for i := 0; i < n; i++ {
		temp = temp.Next
		if temp == nil {
			return head.Next
		}
	}
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		pre = pre.Next
	}

	pre.Next = pre.Next.Next
	return head
}  

