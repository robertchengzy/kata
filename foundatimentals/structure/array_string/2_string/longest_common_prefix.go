package __string

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

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	length := len(strs[0])
	max := -1
loop:
	for i := 0; i < length; i++ {
		cur := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || strs[j][i] != cur {
				break loop
			}
		}
		max = i
	}
	return strs[0][:max+1]
}
