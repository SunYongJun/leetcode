package leetcode

/**
107. 二叉树的层序遍历 II
给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var ans = make([][]int,0)
	var rec func(node *TreeNode, depth int)
	rec = func(node *TreeNode, depth int) {
		if node != nil {
			if len(ans) <= depth {
				ans = append(ans,[]int{})
			}
			rec(node.Left, depth+1)
			rec(node.Right, depth+1)
			ans[depth] = append(ans[depth], node.Val)
		}
	}
	rec(root,0)
	l := len(ans)
	for i := 0 ; i < len(ans) / 2 ; i++ {
		ans[i],ans[l - 1 - i] = ans[l - 1 - i], ans[i]
	}
	return ans
}