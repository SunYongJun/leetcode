package MyLeetCode


/**

	40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

*/


func combinationSum2(candidates []int, target int) [][]int {
	l := len(candidates)
	sort.Slice(candidates,func(i,j int) bool{
		return candidates[i] < candidates[j]
	})

	var dfs func(input []int, offset int, sum int) [][]int
	dfs = func(input []int,offset int, sum int) [][]int {
		if sum > target {
			return [][]int{}
		}

		if sum == target {
			return [][]int{input}
		}

		var ans = make([][]int,0)
		for i := offset; i < l; i++ {
			var tmp = make([]int, 0)
			tmp = append(tmp,input...)
			tmp = append(tmp,candidates[i])
			ans = append(ans, dfs(tmp, i+1, sum + candidates[i] )...)
		}

		return ans
	}
	ans := dfs([]int{}, 0, 0)
	var m = make(map[string][]int)
	for i := range ans {
		num := ""
		for j := range ans[i] {
			num += fmt.Sprintf("%d",ans[i][j])
		}
		m[num] = ans[i]
	}
	var rs = make([][]int,0)
	for i := range m {
		rs = append(rs,m[i])
	}


	return rs
}