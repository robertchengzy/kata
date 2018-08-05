package leetcode

import (
	"math"
	"strings"
)

// Reverse String 反转字符串
/*
	请编写一个函数，其功能是将输入的字符串反转过来。

	示例：

	输入：s = "hello"
	返回："olleh"
*/
func reverseString(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
	return string(r)
}

func reverseString1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}

// Reverse Integer 颠倒整数
/*
	给定一个 32 位有符号整数，将整数中的数字进行反转。

	示例 1:
	输入: 123
	输出: 321
	示例 2:
	输入: -123
	输出: -321
	示例 3:
	输入: 120
	输出: 21
	注意:
	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
*/
// Time Complexity: O(log(x)). There are roughly log10(x) digits in xx.
// Space Complexity: O(1).
func reverseInteger(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}

	return rev
}

func reverseInteger2(x int) int {
	out := 0
	for ; x != 0; x /= 10 {
		out = out*10 + x%10
		if out > math.MaxInt32 || out < math.MinInt32 {
			return 0
		}
	}
	return out
}

// First Unique Character in a String 字符串中的第一个唯一字符
/*
	给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

	案例:

	s = "leetcode"
	返回 0.

	s = "loveleetcode",
	返回 2.


	注意事项：您可以假定该字符串只包含小写字母。
*/
func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	data := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := data[s[i]]; ok {
			continue
		}

		if i == len(s)-1 {
			return len(s) - 1
		}

		repeat := 0
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				repeat += 1
			}

			if j == len(s)-1 && repeat == 0 {
				return i
			}

			data[s[i]] = 1
		}
	}

	return -1
}

func firstUniqChar2(s string) int {
	table := make([]int, 26)
	for _, v := range s {
		table[v-'a']++
	}

	for i, v := range s {
		if table[v-'a'] == 1 {
			return i
		}
	}

	return -1
}

// Valid Anagram 有效的字母异位词
/*
	给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

	示例 1:

	输入: s = "anagram", t = "nagaram"
	输出: true
	示例 2:

	输入: s = "rat", t = "car"
	输出: false
	说明:
	你可以假设字符串只包含小写字母。

	进阶:
	如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/
// Time complexity : O(n). Time complexity is O(n) because accessing the counter table is a constant time operation.
// Space complexity : O(1). Although we do use extra space, the space complexity is O(1) because the table's size stays constant no matter how large nn is.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabet := make([]int, 26)
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++
		alphabet[t[i]-'a']--
	}

	for _, v := range alphabet {
		if v != 0 {
			return false
		}
	}

	return true
}

// Valid Palindrome 验证回文字符串
/*
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

	说明：本题中，我们将空字符串定义为有效的回文串。

	示例 1:

	输入: "A man, a plan, a canal: Panama"
	输出: true
	示例 2:

	输入: "race a car"
	输出: false
*/
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	// build a lowercase/numeric slice from the string
	// Over allocate storage if needed to save re-allocation later
	ln := make([]rune, 0, len(s))

	for _, r := range s {
		rok, ok := fixChar(r)
		if !ok {
			continue
		}
		ln = append(ln, rok)
	}

	// Now check if it's a palindrome
	for i, j := 0, len(ln)-1; i < j; i, j = i+1, j-1 {
		if ln[i] != ln[j] {
			return false
		}
	}

	return true
}

func fixChar(r rune) (okrune rune, ok bool) {
	if ('0' <= r && r <= '9') || ('a' <= r && r <= 'z') {
		return r, true
	}

	if 'A' <= r && r <= 'Z' {
		return r - 'A' + 'a', true
	}

	return '\u0000', false
}

func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}
	ls := strings.ToLower(s)

	for i, j := 0, len(ls)-1; i < j; {
		if !((ls[i] >= '0' && ls[i] <= '9') || (ls[i] >= 'a' && ls[i] <= 'z')) {
			i++
		} else if !((ls[j] >= '0' && ls[j] <= '9') || (ls[j] >= 'a' && ls[j] <= 'z')) {
			j--
		} else {
			if ls[i] != ls[j] {
				return false
			}
			i++
			j--
		}
	}

	return true
}

// String to Integer (atoi) 字符串转整数（atoi）
/*
	实现 atoi，将字符串转为整数。

	在找到第一个非空字符之前，需要移除掉字符串中的空格字符。如果第一个非空字符是正号或负号，选取该符号，并将其与后面尽可能多的连续的数字组合起来，这部分字符即为整数的值。如果第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

	字符串可以在形成整数的字符后面包括多余的字符，这些字符可以被忽略，它们对于函数没有影响。

	当字符串中的第一个非空字符序列不是个有效的整数；或字符串为空；或字符串仅包含空白字符时，则不进行转换。

	若函数不能执行有效的转换，返回 0。

	说明：

	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。如果数值超过可表示的范围，则返回  INT_MAX (2^31 − 1) 或 INT_MIN (−2^31) 。

	示例 1:

	输入: "42"
	输出: 42
	示例 2:

	输入: "   -42"
	输出: -42
	解释: 第一个非空白字符为 '-', 它是一个负号。
		 我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
	示例 3:

	输入: "4193 with words"
	输出: 4193
	解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
	示例 4:

	输入: "words and 987"
	输出: 0
	解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
		 因此无法执行有效的转换。
	示例 5:

	输入: "-91283472332"
	输出: -2147483648
	解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
		 因此返回 INT_MIN (−2^31) 。
*/
func myAtoi(str string) int {
	res, sign, l, idx := 0, 1, len(str), 0

	// Skip leading spaces
	for idx < l && (str[idx] == ' ' || str[idx] == '\t') {
		idx++
	}

	// +/- Sign
	if idx < l {
		if str[idx] == '+' {
			sign = 1
			idx++
		} else if str[idx] == '-' {
			sign = -1
			idx++
		}
	}

	// Numbers
	for idx < l && str[idx] >= '0' && str[idx] <= '9' {
		res = res*10 + int(str[idx]) - '0'
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*res < math.MinInt32 {
			return math.MinInt32
		}
		idx++
	}

	return res * sign
}

// Implement strStr() 实现strStr()
/*
	实现 strStr() 函数。

	给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

	示例 1:

	输入: haystack = "hello", needle = "ll"
	输出: 2
	示例 2:

	输入: haystack = "aaaaa", needle = "bba"
	输出: -1
	说明:

	当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

	对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if haystackLen < needleLen {
		return -1
	}

	for i := 0; i <= haystackLen; i++ {
		if i+needleLen > haystackLen {
			break
		}
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}

// Count and Say 数数并说
/*
	报数序列是指一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

	1.     1
	2.     11
	3.     21
	4.     1211
	5.     111221
	1 被读作  "one 1"  ("一个一") , 即 11。
	11 被读作 "two 1s" ("两个一"）, 即 21。
	21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

	给定一个正整数 n ，输出报数序列的第 n 项。

	注意：整数顺序将表示为一个字符串。

	示例 1:

	输入: 1
	输出: "1"
	示例 2:

	输入: 4
	输出: "1211"
*/
func countAndSay(n int) string {
	return memCountAndSay(n, map[int]string{})
}

func memCountAndSay(n int, m map[int]string) string {
	if n <= 1 {
		return "1"
	}

	if result, ok := m[n]; ok {
		return result
	}

	var result []byte
	prev := countAndSay(n - 1)
	for i := 0; i < len(prev); i++ {
		c := count(prev[i:])
		result = append(result, byte(c+0x30), prev[i])
		i += c - 1
	}

	m[n] = string(result)
	return string(result)
}

func count(s string) int {
	if len(s) == 0 {
		return 0
	}

	first := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != first {
			return i
		}
	}

	return len(s)
}

// Longest Common Prefix 最长公共前缀
/*
	编写一个函数来查找字符串数组中的最长公共前缀。

	如果不存在公共前缀，返回空字符串 ""。

	示例 1:

	输入: ["flower","flow","flight"]
	输出: "fl"
	示例 2:

	输入: ["dog","racecar","car"]
	输出: ""
	解释: 输入不存在公共前缀。
	说明:

	所有输入只包含小写字母 a-z 。
*/
// Approach 1: Horizontal scanning
// Time complexity : O(S), where S is the sum of all characters in all strings.
// In the worst case all nz strings are the same. The algorithm compares the string S1 with the other strings [S2…Sn]
// There are S character comparisons, where S is the sum of all characters in the input array.
// Space complexity : O(1). We only used constant extra space.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}

// Approach 2: Vertical scanning
// Time complexity : O(S) , where S is the sum of all characters in all strings. In the worst case there will be n equal strings with length mm and the algorithm performs S = m*n character comparisons.
// Even though the worst case is still the same as Approach 1, in the best case there are at most n*minLen comparisons where minLenm is the length of the shortest string in the array.
// Space complexity : O(1). We only used constant extra space.
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// Approach 3: Divide and conquer
// Approach 4: Binary search
