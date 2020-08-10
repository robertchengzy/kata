package queue_stack

import (
	"strconv"
	"strings"
)

/*
字符串解码
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
示例 1：
	输入：s = "3[a]2[bc]"
	输出："aaabcbc"
示例 2：
	输入：s = "3[a2[c]]"
	输出："accaccacc"
示例 3：
	输入：s = "2[abc]3[cd]ef"
	输出："abcabccdcdcdef"
示例 4：
	输入：s = "abc3[cd]xyz"
	输出："abccdcdcdxyz"
*/

func decodeString(s string) string {
	var stack []string
	var ptr int
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			stack = append(stack, getDigits(s, &ptr))
		} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' || cur == '[' {
			stack = append(stack, string(cur))
			ptr++
		} else {
			ptr++
			var data []string
			for stack[len(stack)-1] != "[" {
				data = append(data, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			length := len(data)
			for i := 0; i < length/2; i++ {
				data[i], data[length-i-1] = data[length-i-1], data[i]
			}
			stack = stack[:len(stack)-1]
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(getString(data), repTime))
		}
	}
	return getString(stack)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
