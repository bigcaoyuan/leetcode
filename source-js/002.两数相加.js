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
var addTwoNumbers = function (l1, l2) {
	// 记录进位
	let decade = 0,
		v1, v2,
		sum,
		result = new ListNode(0),
		currentNode = result

	while (l1 || l2) {
		v1 = l1 ? l1.val : 0

		v2 = l2 ? l2.val : 0

		sum = v1 + v2 + decade

		decade = sum >= 10 ? 1 : 0
		sum = sum >= 10 ? sum - 10 : sum

		const node = new ListNode(sum)

		currentNode.next = node
		currentNode = currentNode.next

		l1 = l1 ? l1.next : null
		l2 = l2 ? l2.next : null
	}

	// 两个链表都遍历完了，但是还有进位
	if (decade > 0) {
		currentNode.next = new ListNode(decade)
	}

	return result.next
};