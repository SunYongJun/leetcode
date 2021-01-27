package leetcode

/**
332. 重新安排行程
给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。



提示：

如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
所有的机场都用三个大写字母表示（机场代码）。
假定所有机票至少存在一种合理的行程。
所有的机票必须都用一次 且 只能用一次。


示例 1：

输入：[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
输出：["JFK", "MUC", "LHR", "SFO", "SJC"]
示例 2：

输入：[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。

*/

func findItinerary(tickets [][]string) []string {
	l := len(tickets)
	rs := rec("JFK",tickets,[]string{"JFK"})

	fnl := -1
	for i := 0 ; i < len(rs); i++ {
		if len(rs[i]) < l {
			continue;
		}
		if fnl == -1 {
			fnl = i
		}else{
			for j := 0 ; j < l ; j++ {
				if rs[i][j] == rs[fnl][j] {
					continue
				}
				if rs[i][j] < rs[fnl][j] {
					fnl = i
				}
				break
			}
		}
	}
	return rs[fnl]
}

func rec(f string, source [][]string, anses []string) [][]string {

	rs := make([][]string,0)

	for i := 0 ; i < len(source); i++ {
		if source[i][0] == f {
			nf := source[i][1]
			ns := make([][]string,0)
			ns = append(ns, source[:i]...)
			ns = append(ns, source[i+1:]...)

			ans := append(anses,nf)

			if len(ns) == 0 {
				rs = append(rs, ans)
			}else{
				nans := rec(nf, ns, ans )
				rs = append(rs,nans...)
			}


		}
	}

	return rs
}