package MyLeetCode

/**
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/


var m = map[byte]string{
	'1':"",
	'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"pqrs",
	'8':"tuv",
	'9':"wxyz",
}
func letterCombinations(digits string) []string {
	return rec(digits,[]string{})
}

func rec(digits string,rs []string)[]string{
	if len(digits) == 0 {
		return rs
	}
	var ans = make([]string,0)

	s := digits[0]
	tmp := m[s]

	if len(rs) == 0 {
		for i := 0 ; i < len(tmp); i++ {
			ans = append(ans,string(tmp[i]))
		}
	}else{
		for i := 0 ; i < len(tmp); i++ {
			for j := 0 ; j < len(rs); j++ {
				ans = append(ans,rs[j] + string(tmp[i]))
			}

		}
	}

	return rec(digits[1:len(digits)],ans)

}