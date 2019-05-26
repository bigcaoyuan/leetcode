/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] k个一组翻转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	stack := make([]*ListNode, k)
	npre := &ListNode{0, head}
	pre := npre
	temp := head
	cnt := 0
	for {
		for {
			if temp != nil && cnt != k {
				stack[cnt] = temp
				temp = temp.Next
				cnt++
			} else {
				break
			}
		}

		if cnt == k {
			for {
				if cnt == 0 {
					break
				}
				pre.Next = stack[cnt-1]
				pre = pre.Next
				cnt--
			}
			pre.Next = temp
		} else {
			break
		}
	}
	return npre.Next
}

