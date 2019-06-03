/**
 * @param {number} dividend
 * @param {number} divisor
 * @return {number}
 */
var divide = function (dividend, divisor) {
	const SAFE_BIT_MIN = -1073741824
	const SAFE_BIT_MAX = 1073741823
	const INTERGER_MAX = Math.pow(2, 31) - 1
	const INTERGER_MIN = -Math.pow(2, 31)

	let quotient = 0
		, originDevisor = divisor
		, t = 1

	if (dividend >= 0 && divisor > 0 || dividend <= 0 && divisor < 0) {
		while (Math.abs(dividend) >= Math.abs(divisor) || t > 1) {
			if (dividend === 0 || quotient >= INTERGER_MAX)
				break

			if (Math.abs(dividend) < Math.abs(divisor) && Math.abs(divisor >> 1) >= Math.abs(originDevisor)) {
				// 回滚divisor、t的值
				divisor = divisor >> 1
				t = t >> 1
			} else {
				dividend -= divisor
				quotient += t

				if ((divisor <= SAFE_BIT_MAX && divisor >= SAFE_BIT_MIN) && (t <= SAFE_BIT_MAX && t >= SAFE_BIT_MIN && t >= 1)) {
					divisor = divisor << 1
					t = t << 1
				}
			}
		}
	} else {
		while (Math.abs(dividend) >= Math.abs(divisor) || t > 1) {
			if (dividend === 0 || quotient <= INTERGER_MIN)
				break
			if (Math.abs(dividend) < Math.abs(divisor) && Math.abs(divisor >> 1) >= Math.abs(originDevisor)) {
				divisor = divisor >> 1
				t = t >> 1
			} else {
				dividend += divisor
				quotient -= t

				if ((divisor <= SAFE_BIT_MAX && divisor >= SAFE_BIT_MIN) && (t <= SAFE_BIT_MAX && t >= SAFE_BIT_MIN && t >= 1)) {
					divisor = divisor << 1
					t = t << 1
				}
			}
		}
	}

	if (quotient > INTERGER_MAX || quotient < INTERGER_MIN) {
		quotient = INTERGER_MAX
	}

	return quotient
};