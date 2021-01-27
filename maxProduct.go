package leetcode

/**
152. 乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。



示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/

func max(x, y int) int{
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int{
	if x < y {
		return x
	}
	return y
}

func maxProduct(nums []int) int {
	maxVal := nums[0]
	var a,b int = 1,1

	for i:= 0; i< len(nums);i++ {
		aa := nums[i] * a
		bb := nums[i] * b

		a = min(nums[i], min(aa, bb))
		b = max(nums[i], max(aa, bb))

		maxVal = max(maxVal, b)
	}
	return maxVal
}