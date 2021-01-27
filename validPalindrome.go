package MyLeetCode

/**
680. 验证回文字符串 Ⅱ
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:

字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
*/

func validPalindrome(s string) bool {
	j := len(s) - 1
	i := 0


	for ;i < j; {
		if s[i] != s[j] {
			return isValid(s[i+1:j+1]) || isValid(s[i:j])
		}
		i++
		j--
	}
	return true

}

func isValid(s string) bool{
	j := len(s) - 1
	i := 0


	for ;i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true

}