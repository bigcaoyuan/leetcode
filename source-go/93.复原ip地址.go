import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

func restoreIpAddresses(s string) []string {
	strs := strings.Split(s, "")
	res := []string{}
	doRestoreIpAddresses(&strs, 0, 0, "", &res)
	return res
}

func doRestoreIpAddresses(strs *[]string, index int, cnt int, s string, res *[]string) {
	if index == len(*strs) || cnt == 4 {
		if index == len(*strs) && cnt == 4 {
			*res = append(*res, s[1:])
		}
		return
	}

	if cnt == 3 && len(*strs)-index > 3 {
		return
	}

	if (*strs)[index] == "0" {
		doRestoreIpAddresses(strs, index+1, cnt+1, s+".0", res)
		return
	}

	doRestoreIpAddresses(strs, index+1, cnt+1, s+"."+(*strs)[index], res)
	if index <= len(*strs)-2 {
		doRestoreIpAddresses(strs, index+2, cnt+1, s+"."+(*strs)[index]+(*strs)[index+1], res)
	}
	if index <= len(*strs)-3 {
		if (*strs)[index] == "1" {
			doRestoreIpAddresses(strs, index+3, cnt+1, s+"."+(*strs)[index]+(*strs)[index+1]+(*strs)[index+2], res)
		} else if (*strs)[index] == "2" {
			if (*strs)[index+1] <= "4" {
				doRestoreIpAddresses(strs, index+3, cnt+1, s+"."+(*strs)[index]+(*strs)[index+1]+(*strs)[index+2], res)
			} else if (*strs)[index+1] == "5" && (*strs)[index+2] <= "5" {
				doRestoreIpAddresses(strs, index+3, cnt+1, s+"."+(*strs)[index]+(*strs)[index+1]+(*strs)[index+2], res)
			}
		}
	}
}


