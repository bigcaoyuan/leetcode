/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	root := initPrefixNode(strs[0])
	if root == nil {
		return ""
	}
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		node := root
		pre := root
		for _, j := range strs[i] {
			if node == nil {
				break
			} else {
				if node.value == j {
					pre = node
					node = node.next
				} else {
					if node == root {
						return ""
					}
					pre.next = nil
					break
				}
			}
		}
		pre.next = nil
	}
	str := ""
	for {
		if root == nil {
			break
		}
		str = str + string(root.value)
		root = root.next
	}
	return str
}

type PrefixNode struct {
	value rune
	next  *PrefixNode
}

func initPrefixNode(str string) *PrefixNode {
	if len(str) == 0 {
		return nil
	}
	root := &PrefixNode{}
	node := root
	pre := root
	for _, val := range str {
		node.value = val
		node.next = &PrefixNode{}
		pre = node
		node = node.next
	}
	pre.next = nil
	return root
}


