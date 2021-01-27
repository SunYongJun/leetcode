package MyLeetCode

/**
	51. N 皇后
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。



示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]


提示：

1 <= n <= 9
皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
*/

func solveNQueens(n int) [][]string {
	var rec func (ans []int, row int) [][]int

	rec = func (ans []int, row int) [][]int{

		if row == n {
			return [][]int{ans}
		}
		var rs = make([][]int,0)

		for i := 0 ; i < n ; i++{
			v := 1 << i
			flag := true
			for k := range ans {

				if ans[k] == v || (ans[k] << (row - k)) == v || (ans[k] >> (row - k)) == v {
					flag = false
					break
				}
			}

			if flag {
				var cp = make([]int,0)
				cp = append(cp, ans...)
				cp = append(cp, v)
				rs = append(rs,rec( cp, row + 1)...)
			}
		}

		return rs
	}
	ints := rec([]int{},0)

	var strs = make([][]string,0)
	for i := 0 ; i < len(ints); i++ {

		var strList = make([]string,0)
		for j := 0 ; j < n ; j++ {
			item := ""
			flag := n + 1

			for k := 1 ; k <= n ; k++{

				if k > flag || (ints[i][j] >> k) > 0 {
					item += "."
				}else{
					item += "Q"
					flag = k
				}
			}
			strList = append(strList, item)
		}
		strs = append(strs, strList)
	}
	return strs
}
