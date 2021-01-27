package leetcode

/**
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/

func searchInsert(nums []int, target int) int {
	var dfs func(start int, end int) int
	dfs = func(start int, end int) int {

		if nums[start]  > target {
			return start
		}
		if nums[end] < target {
			return end + 1
		}

		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			return dfs(start,mid )
		}else{
			return dfs(mid + 1,end)
		}
	}
	return dfs(0,len(nums)-1)
}