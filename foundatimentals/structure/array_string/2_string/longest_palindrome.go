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

/*
动态规划
时间复杂度：O(n^2)，其中 n 是字符串的长度。动态规划的状态总数为 O(n^2)，对于每个状态，我们需要转移的时间为 O(1)。
空间复杂度：O(n^2)，即存储动态规划状态需要的空间。
*/
func longestPalindrome1(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

/*
中心扩展算法
时间复杂度：O(n^2)，其中 n 是字符串的长度。长度为 1 和 2 的回文中心分别有 n 和 n-1 个，每个回文中心最多会向外扩展 O(n) 次。
空间复杂度：O(1)。
*/
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

/*
Manacher 算法
时间复杂度：O(n)，其中 n 是字符串的长度。由于对于每个位置，扩展要么从当前的最右侧臂长 right 开始，要么只会进行一步，而 right 最多向前走 O(n) 步，因此算法的复杂度为 O(n)。
空间复杂度：O(n)，我们需要 O(n) 的空间记录每个位置的臂长。
*/
func longestPalindrome3(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	s = t
	arm_len := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var cur_arm_len int
		if right >= i {
			i_sym := j*2 - i
			min_arm_len := min(arm_len[i_sym], right-i)
			cur_arm_len = expand(s, i-min_arm_len, i+min_arm_len)
		} else {
			cur_arm_len = expand(s, i, i)
		}
		arm_len = append(arm_len, cur_arm_len)
		if i+cur_arm_len > right {
			j = i
			right = i + cur_arm_len
		}
		if cur_arm_len*2+1 > end-start {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}
