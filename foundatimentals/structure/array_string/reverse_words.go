package array_string

import "strings"

/*
反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
示例 1:
	输入: "Let's take LeetCode contest"
	输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
*/

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	l := len(words)
	var res []string
	for i := 0; i < l; i++ {
		data := []byte(words[i])
		n := len(data)
		for j := 0; j < n/2; j++ {
			data[j], data[n-1-j] = data[n-1-j], data[j]
		}
		res = append(res, string(data))
	}
	return strings.Join(res, " ")
}
