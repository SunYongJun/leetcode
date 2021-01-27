package leetcode


/**

109. 有序链表转换二叉搜索树
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedListToBST(head *ListNode) *TreeNode {

	return treeBuild(head,nil)
}

func getMid(left,right *ListNode) *ListNode{
	s,l := left, left
	for s != right && s.Next != right {
		s = s.Next.Next
		l = l.Next
	}

	return l
}

func treeBuild(left,right *ListNode) *TreeNode{
	if left == right {
		return nil
	}
	mid := getMid(left,right)
	root := &TreeNode{
		Val: mid.Val,

	}

	root.Left = treeBuild(left,mid)
	root.Right = treeBuild(mid.Next,right)
	return root
}