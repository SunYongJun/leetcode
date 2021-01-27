package MyLeetCode

/**

214. 最短回文串
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。



示例 1：

输入：s = "aacecaaa"
输出："aaacecaaa"
示例 2：

输入：s = "abcd"
输出："dcbabcd"


提示：

0 <= s.length <= 5 * 104
s 仅由小写英文字母组成

*/

func shortestPalindrome(s string) string {
	l := len(s)
	pre := ""
	pi := 0

	for i := 0 ; i < l ; i++{
		pre = string(s[i]) + pre
		if s[:i+1] == pre{
			pi = i
		}
	}

	f := ""
	for j := pi + 1 ; j < l ; j++{
		f = string(s[j]) + f
	}
	return f + s
}