package string

/*
最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。
示例 2：
	输入: "cbbd"
	输出: "bb"
*/

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var start, end, res = 0, 1, 1
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				var ok = true
				for p1, p2 := i+1, j-1; p1 < p2; p1, p2 = p1+1, p2-1 {
					if s[p1] != s[p2] {
						ok = false
						break
					}
				}

				if ok && res < j-i+1 {
					res = j - i + 1
					start, end = i, j+1
				}
			}
		}
	}
	return s[start:end]
}
