/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function (n) {
	let ans = []
		, backtracking = function (cur, open, close) {
			if (cur.length === n * 2) {
				ans.push(cur)
			}

			if (open < n)
				backtracking(cur + '(', open + 1, close)
			if (close < open)
				backtracking(cur + ')', open, close + 1)
		}

	backtracking('', 0, 0)

	return ans
};