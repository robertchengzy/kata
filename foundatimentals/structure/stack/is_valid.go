package stack

/*
有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
示例 1:
输入: "()"
输出: true
示例 2:
输入: "()[]{}"
输出: true
示例 3:
输入: "(]"
输出: false
示例 4:
输入: "([)]"
输出: false
示例 5:
输入: "{[]}"
输出: true
*/

func isValid(s string) bool {
	pair := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)
	for _, v := range []byte(s) {
		if _, ok := pair[v]; ok {
			stack = append(stack, v)
		} else {
			stackLen := len(stack)
			if stackLen > 0 && pair[stack[stackLen-1]] == v {
				stack = stack[:stackLen-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
