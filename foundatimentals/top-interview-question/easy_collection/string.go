package leetcode

import (
	"math"
	"strings"
)

// Reverse String
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

// Reverse Integer
// Time Complexity: O(log(x)). There are roughly log10(x) digits in xx.
// Space Complexity: O(1).
func reverseInteger(x int) int {
	rev := 0
	for {
		if x == 0 {
			break
		}

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

// First Unique Character in a String
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

// Valid Anagram
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

// Valid Palindrome
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

// String to Integer (atoi)
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

// Implement strStr()
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

// Count and Say
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

// Longest Common Prefix
// Time complexity : O(S), where S is the sum of all characters in all strings.
// In the worst case all nz strings are the same. The algorithm compares the string S1 with the other strings [S2…Sn]
// There are S character comparisons, where S is the sum of all characters in the input array.
// Space complexity : O(1). We only used constant extra space.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}

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
				return strs[0][0:i]
			}
		}
	}

	return strs[0]
}

// TODO Divide and conquer && Binary search etc.
