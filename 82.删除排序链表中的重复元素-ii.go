/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	npre := &ListNode{head.Val - 1, head}
	pre := npre
	cur := head
	for {
		if cur == nil {
			break
		}
		for {
			if cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			} else {
				break
			}
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return npre.Next
}

