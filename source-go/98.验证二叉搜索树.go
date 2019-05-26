/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	findLastVal := false
	lastValue := 0
	res := []*TreeNode{}
	for {
		for {
			if root != nil {
				res = append(res, root)
				root = root.Left
				continue
			}
			break
		}
		if len(res) == 0 {
			break
		}
		root = res[len(res)-1]
		res = res[:len(res)-1]
		if !findLastVal {
			findLastVal = true
			lastValue = root.Val
		} else {
			if lastValue >= root.Val {
				return false
			}
			lastValue = root.Val
		}
		root = root.Right
	}
	return true
}

