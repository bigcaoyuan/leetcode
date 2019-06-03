/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
  let same = true, ret = ''

  if (!strs[0]) return ''

  for (let i = 0; i < strs[0].length; i++) {
    for (let j = 1; j < strs.length; j++) {
      if (strs[0][i] !== strs[j][i]) {
        same = false
        break
      }
    }
    if (!same) {
      break
    } else {
      ret += strs[0][i]
    }
  }

  return ret
};