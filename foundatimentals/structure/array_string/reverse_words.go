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

/*
思路与算法
	开辟一个新字符串。然后从头到尾遍历原字符串，直到找到空格为止，此时找到了一个单词，并能得到单词的起止位置。随后，根据单词的起止位置，
	可以将该单词逆序放到新字符串当中。如此循环多次，直到遍历完原字符串，就能得到翻转后的结果。
时间复杂度：O(N)，其中 N 为字符串的长度。原字符串中的每个字符都会在 O(1) 的时间内放入新字符串中。
空间复杂度：O(N)。我们开辟了与原字符串等大的空间。
*/
func reverseWords1(s string) string {
	length := len(s)
	var ret []byte
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := i - 1; p >= start; p-- {
			ret = append(ret, s[p])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}

/*
思路与算法
	此题也可以直接在原字符串上进行操作，避免额外的空间开销。当找到一个单词的时候，我们交换字符串第一个字符与倒数第一个字符，
	随后交换第二个字符与倒数第二个字符…… 如此反复，就可以在原空间上翻转单词。
时间复杂度：O(N)。字符串中的每个字符要么在 O(1) 的时间内被交换到相应的位置，要么因为是空格而保持不动。
空间复杂度：O(1)。因为不需要开辟额外的数组。
*/
func reverseWords2(s string) string {
	b := []byte(s)
	i, length := 0, len(s)
	for i < length {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		left, right := start, i-1
		for left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
		for i < length && s[i] == ' ' {
			i++
		}
	}
	return string(b)
}
