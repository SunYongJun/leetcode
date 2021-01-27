package MyLeetCode

/**

530. 二叉搜索树的最小绝对差
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。



示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。


提示：

树中至少有 2 个节点。
本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {

	ans,r := math.MaxInt64,-1

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {

			dfs(node.Left)

			if r != -1 && node.Val - r < ans {
				ans = node.Val - r
			}

			r = node.Val

			dfs(node.Right)
		}
	}
	dfs(root)

	return ans
}