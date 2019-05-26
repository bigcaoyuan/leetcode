/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for {
		for {
			if root == nil {
				break
			}
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) == 0 {
			break
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// 递归思路
// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	doInorderTraversal(root, &res)
// 	return res
// }

// func doInorderTraversal(root *TreeNode, res *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	doInorderTraversal(root.Left, res)
// 	*res = append(*res, root.Val)
// 	doInorderTraversal(root.Right, res)
// }

