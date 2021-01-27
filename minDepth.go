package MyLeetCode

/**

111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5


提示：

树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	depth := 0
	if root != nil {
		depth++

		if root.Left != nil && root.Right == nil {
			depth += minDepth(root.Left)
		}else if root.Left == nil && root.Right != nil  {
			depth +=  minDepth(root.Right)
		}else if root.Left != nil && root.Right != nil {
			depth += min( minDepth(root.Left), minDepth(root.Right))
		}
	}else{
		return 0
	}


	return depth
}

func min(a,b int) int{
	if a > b {
		return b
	}else{
		return a
	}
}