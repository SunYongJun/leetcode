package leetcode


/**
687. 最长同值路径
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:

2
示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:

2
注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。

*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	var ans int
	var longes func (node *TreeNode) int
	longes = func (node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftLong := longes(node.Left)
		rightLong := longes(node.Right)
		var l,r int
		if node.Left != nil && node.Left.Val == node.Val {
			l = leftLong + 1
		}

		if node.Right != nil && node.Right.Val == node.Val{
			r = rightLong + 1
		}

		ans = max(ans,l + r)
		return max(l,r)
	}
	longes(root)
	return ans
}

func max(a,b int) int{
	if a > b {
		return a
	}
	return b
}
