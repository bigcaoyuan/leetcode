/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function (l1, l2) {
	let current1 = l1
		, current2 = l2
		, dummy = new ListNode(0)
		, node = dummy
	while (current1 && current2) {
		if (current1.val > current2.val) {
			node.next = new ListNode(current2.val)
			current2 = current2.next
		} else {
			node.next = new ListNode(current1.val)
			current1 = current1.next
		}
		node = node.next
	}

	while (current1) {
		node.next = new ListNode(current1.val)
		current1 = current1.next
		node = node.next
	}

	while (current2) {
		node.next = new ListNode(current2.val)
		current2 = current2.next
		node = node.next
	}

	return dummy.next
};