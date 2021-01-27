package MyLeetCode

/**
199. 二叉树的右视图
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {

	ans := make([]int,0)
	list := make([]*TreeNode,0)
	if root != nil {
		list = append(list,root)
	}
	for len(list) > 0 {
		//把每层的节点从左到右加入tmp，然后赋值给list
		tmp := make([]*TreeNode,0)
		for i := range list {
			if list[i].Left != nil {
				tmp = append(tmp, list[i].Left)
			}
			if list[i].Right != nil {
				tmp = append(tmp, list[i].Right)
			}

		}
		//取list最后一个节点的值，追加到slice
		ans = append(ans, list[len(list) - 1].Val)
		list = tmp
	}
	return ans
}