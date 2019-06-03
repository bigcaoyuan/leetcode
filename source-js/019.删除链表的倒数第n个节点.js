/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function (head, n) {
	let dummyNode = new ListNode(0)
		, first = dummyNode
		, second = dummyNode

	dummyNode.next = head

	for (let i = 0; i <= n; i++) {
		first = first.next
	}

	while (first) {
		first = first.next
		second = second.next
	}

	second.next = second.next.next

	return dummyNode.next
};