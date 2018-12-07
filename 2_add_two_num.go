/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2 int
	res := &ListNode{}
	next := res
	plus := false
	for {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}

		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		if plus == true {
			v1++
			plus = false
		}

		if v1+v2 >= 10 {
			next.Val = v1 + v2 - 10
			plus = true
		} else {
			next.Val = v1 + v2
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			break
		}
		next_ := &ListNode{}
		next.Next = next_
		next = next_
	}
	return res
}