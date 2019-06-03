/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
	let stack = []

	for (let i = 0; i < s.length; i++) {
		switch (s[i]) {
			case '(':
				stack.push('(')
				break
			case ')':
				if (stack.length === 0 || stack.pop() !== '(')
					return false
				break
			case '[':
				stack.push('[')
				break
			case ']':
				if (stack.length === 0 || stack.pop() !== '[')
					return false
				break
			case '{':
				stack.push('{')
				break
			case '}':
				if (stack.length === 0 || stack.pop() !== '{')
					return false
				break
			default:
				return false

		}
	}

	return stack.length === 0
};