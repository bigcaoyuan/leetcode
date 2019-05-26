/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	last := lists[0]
	npre := &ListNode{}
	npre.Next = last
	temp := npre
	for i := 1; i < len(lists); i++ {
		root1 := last
		root2 := lists[i]
		for {
			if root1 == nil {
				temp.Next = root2
				break
			} else if root2 == nil {
				temp.Next = root1
				break
			}
			if root1.Val < root2.Val {
				temp.Next = root1
				temp = temp.Next
				root1 = root1.Next
			} else {
				temp.Next = root2
				temp = temp.Next
				root2 = root2.Next
			}
		}
		last = npre.Next
		temp = npre
	}
	return npre.Next
}

// 方法太low了
func mergeKLists1(lists []*ListNode) *ListNode {
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
			i--
		}
	}
	npre := &ListNode{}
	temp := npre
	for {
		if len(lists) == 0 {
			break
		}
		min := lists[0].Val
		index := 0
		for i := 1; i < len(lists); i++ {
			if lists[i].Val < min {
				index = i
				min = lists[i].Val
			}
		}
		temp.Next = lists[index]
		lists[index] = lists[index].Next
		temp = temp.Next
		if lists[index] == nil {
			lists = append(lists[:index], lists[index+1:]...)
		}
	}
	return npre.Next
}



