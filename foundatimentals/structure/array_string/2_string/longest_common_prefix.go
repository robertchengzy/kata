package string

/*
最长公共前缀
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

/*
横向扫描
时间复杂度：O(mn)，其中 m 是字符串数组中的字符串的平均长度，n 是字符串的数量。最坏情况下，字符串数组中的每个字符串的每个字符都会被比较一次。
空间复杂度：O(1)。使用的额外空间复杂度为常数。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
纵向扫描
时间复杂度：O(mn)，其中 m 是字符串数组中的字符串的平均长度，n 是字符串的数量。最坏情况下，字符串数组中的每个字符串的每个字符都会被比较一次。
空间复杂度：O(1)。使用的额外空间复杂度为常数。
*/
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

/*
分治
时间复杂度：O(mn)，其中 m 是字符串数组中的字符串的平均长度，n 是字符串的数量。时间复杂度的递推式是 T(n)=2⋅T(2/n)+O(m)，通过计算可得 T(n)=O(mn)。
空间复杂度：O(mlogn)，其中 m 是字符串数组中的字符串的平均长度，n 是字符串的数量。空间复杂度主要取决于递归调用的层数，层数最大为 logn，每层需要 m 的空间存储返回结果。
*/
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

/*
二分查找
时间复杂度：O(mnlogm)，其中 m 是字符串数组中的字符串的最小长度，n 是字符串的数量。二分查找的迭代执行次数是 O(logm)，每次迭代最多需要比较 mn 个字符，因此总时间复杂度是 O(mnlogm)。
空间复杂度：O(1)。使用的额外空间复杂度为常数。
*/
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
