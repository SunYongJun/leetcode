package leetcode

/**

39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/


func combinationSum(candidates []int, target int) [][]int {
	var rec func(input []int,offset int) [][]int

	rec = func(input []int,offset int) [][]int{
		num := 0
		for i := range input{
			num += input[i]
		}
		if num > target{
			return [][]int{}
		}
		if num == target{
			return [][]int{input}
		}
		l := len(input)
		var ans = make([][]int,0)
		for i := offset ; i < len(candidates); i++  {
			var tmp = make([]int,l)
			copy(tmp,input)
			tmp = append(tmp,candidates[i])
			ans = append(ans, rec(tmp,i)...)
		}
		return ans
	}
	return rec([]int{},0)
}