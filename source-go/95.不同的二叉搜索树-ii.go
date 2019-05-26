/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return doGenerateTrees(1, n)
}

func doGenerateTrees(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	if start == end {
		return []*TreeNode{&TreeNode{start, nil, nil}}
	}
	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := doGenerateTrees(start, i-1)
		right := doGenerateTrees(i+1, end)
		for m := 0; m < len(left); m++ {
			for n := 0; n < len(right); n++ {
				res = append(res, &TreeNode{i, left[m], right[n]})
			}
		}
	}
	return res
}

