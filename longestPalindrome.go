package leetcode

/**
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成
*/

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	if len(s) == 1 {
		return string(s[0])
	}
	t := "#"

	for i:=0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}

	length := len(t)

	dp := make([]int,length)

	max,middle := 0,0

	for i:=0 ; i < length; i++ {
		dp[i] = 0
		for j := i ; j > i / 2 ; j-- {
			if dp[j] >= 0 {
				if  t[i] == t[2 * j - i] {
					dp[j] += 1
				}else {

					if dp[j] > max {
						max = dp[j]
						middle =  j
					}

					dp[j] = -1
				}
			}

		}
	}

	if max == 1 {
		return string(s[0])
	}

	s = t[middle-max+1:middle+max]
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans

}