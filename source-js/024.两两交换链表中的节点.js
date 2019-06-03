/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function (head) {
	let dummy = new ListNode(0)
		, prev = dummy
		, current = head
		, next = null

	dummy.next = head

	while (current && current.next) {
		next = current.next

		prev.next = next
		current.next = next.next
		next.next = current

		prev = current
		current = current.next
	}

	return dummy.next
};