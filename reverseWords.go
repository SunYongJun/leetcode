package MyLeetCode

/**
557. 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。



示例：

输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"


提示：

在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
*/

func reverseWords(s string) string {
	l := len(s)
	c := []byte(s)
	for i := 0 ; i * 2 < l  ; i++ {
		c[i], c[l - 1 - i] = c[l - 1 -i],c[i]
	}

	ans := ""
	flag :=  l
	for j := l - 1 ; j >= 0 ; j-- {
		if c[j] == ' '{
			ans += string(c[j+1:flag]) + " "
			flag = j
		}
		if j == 0 {
			ans += string(c[j:flag + 1])
		}
	}
	return ans[:l]
}