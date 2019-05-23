/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	lastVal := head.Val
	pre := head
	temp := head.Next
	for {
		if temp == nil {
			break
		}
		if temp.Val != lastVal {
			lastVal = temp.Val
			pre = temp
		} else {
			pre.Next = temp.Next
		}
		temp = temp.Next
	}
	return head
}

