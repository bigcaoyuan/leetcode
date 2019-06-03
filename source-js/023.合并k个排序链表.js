/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
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
/**
* @param {ListNode[]} lists
* @return {ListNode}
*/
var mergeKLists = function (lists) {
	if (lists.length === 0) {
		return null
	} else if (lists.length === 1) {
		return lists[0]
	} else {
		return mergeTwoLists(mergeKLists(lists.slice(0, Math.floor(lists.length / 2))), mergeKLists(lists.slice(Math.floor(lists.length / 2), lists.length)))
	}
};