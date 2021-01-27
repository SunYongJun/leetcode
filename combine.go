package leetcode

/**
77. 组合
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {

	var rec func(input []int) [][]int
	rec = func(input []int) [][]int{
		if len(input) == k {
			return [][]int{input}
		}

		var ans = make([][]int,0)
		start := 1
		l := len(input)
		if len(input) > 0 {
			start = input[len(input) - 1] + 1
		}
		for i := start ; i <= n; i++ {
			p := make([]int,l)
			copy(p,input)
			p = append(p,i)

			anses := rec(p)
			for j := range anses {
				ans = append(ans, anses[j])
			}

		}
		return ans
	}
	return rec([]int{})
}