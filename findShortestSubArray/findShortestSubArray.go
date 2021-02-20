package findShortestSubArray

/*
697. 数组的度
给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。



示例 1：

输入：[1, 2, 2, 3, 1]
输出：2
解释：
输入数组的度是2，因为元素1和2的出现频数最大，均为2.
连续子数组里面拥有相同度的有如下所示:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组[2, 2]的长度为2，所以返回2.
示例 2：

输入：[1,2,2,3,1,4,2]
输出：6


提示：

nums.length 在1到 50,000 区间范围内。
nums[i] 是一个在 0 到 49,999 范围内的整数。
*/

func FindShortestSubArray(nums []int) int {

	m := make(map[int]int)
	arr := make([][]int,1)

	for i := 0 ; i < len(nums); i++ {
		if _,ok := m[nums[i]]; ok {
			m[nums[i]]++
			if len(arr) < m[nums[i]] {
				arr = append(arr, []int{nums[i]})
			}else{
				arr[m[nums[i]] - 1] = append(arr[m[nums[i]] - 1],nums[i])
			}
		}else{
			m[nums[i]] = 1
			arr[0] = append(arr[0], nums[i])
		}
	}

	l := len(arr)
	max := arr[l - 1]

	ans := len(nums)

	for i := 0 ; i < len(max); i++ {
		start,end := 0,len(nums)-1
		for j := 0 ; j < len(nums); j++ {
			if nums[j] == max[i] {
				start = j
				break
			}
		}

		for j := len(nums) - 1; j > 0 ; j-- {
			if nums[j] == max[i] {
				end = j
				break
			}
		}

		diff := end - start + 1
		if ans > diff {
			ans = diff
		}
	}


	return ans
}