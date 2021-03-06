package string

import "strings"

/*
翻转字符串里的单词
给定一个字符串，逐个翻转字符串中的每个单词。
示例 1：
	输入: "the sky is blue"
	输出: "blue is sky the"
示例 2：
	输入: "  hello world!  "
	输出: "world! hello"
	解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：
	输入: "a good   example"
	输出: "example good a"
	解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
说明：
	无空格字符构成一个单词。
	输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
	如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
进阶：
	请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。
*/

/*
时间复杂度：O(N)，其中 N 为输入字符串的长度。
空间复杂度：O(N)，用来存储字符串分割之后的结果。
*/
func reverseWords(s string) string {
	strs := strings.Split(strings.TrimSpace(s), " ")
	var res []string
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != "" {
			res = append(res, strs[i])
		}
	}
	return strings.Join(res, " ")
}

/*
双端队列
时间复杂度：O(N)，其中 N 为输入字符串的长度。
空间复杂度：O(N)，双端队列存储单词需要 O(N) 的空间。
*/
func reverseWords1(s string) string {
	left, right := 0, len(s)-1
	for left <= right && s[left] == ' ' {
		left++
	}
	for left <= right && s[right] == ' ' {
		right--
	}
	var deque []string
	var word string
	for left <= right {
		c := s[left]
		if len(word) != 0 && c == ' ' {
			deque = append([]string{word}, deque...)
			word = ""
		} else if c != ' ' {
			word += string(c)
		}
		left++
	}
	deque = append([]string{word}, deque...)

	return strings.Join(deque, " ")
}
