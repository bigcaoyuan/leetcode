/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */
//  "What   must   be","acknowledgment  ","shall be         "
//  "What   must   be","acknowledgment  ","shall be        "

// This    is    an","example  of text","justification.   "
// This    is    an","example  of text","justification.  "
func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	tmp := []string{}
	cnt := 0
	width := 0
	for i := 0; i < len(words); {
		if cnt == 0 || width+cnt+len(words[i]) <= maxWidth {
			tmp = append(tmp, words[i])
			cnt++
			width += len(words[i])
			i++
		} else {
			black := maxWidth - width
			if cnt == 1 {
				str := tmp[0]
				for j := 0; j < black; j++ {
					str += " "
				}
				res = append(res, str)
			} else {
				equal := black / (cnt - 1)
				over := (black - equal*(cnt-1)) % (cnt - 1)
				str := ""
				for j := 0; j < cnt; j++ {
					str += tmp[j]
					if j == cnt-1 {
						break
					}
					for m := 0; m < equal; m++ {
						str += " "
					}
					if j < over {
						str += " "
					}
				}
				res = append(res, str)
			}
			tmp = []string{}
			cnt = 0
			width = 0
		}
	}
	str := ""
	for i := 0; i < len(tmp); i++ {
		str += tmp[i]
		if i != len(tmp)-1 {
			str += " "
		}
	}
	for i := 0; i <= maxWidth-width-cnt; i++ {
		str += " "
	}
	res = append(res, str)
	return res
}

