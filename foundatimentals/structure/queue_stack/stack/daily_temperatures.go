package stack

/*
每日温度
请根据每日气温列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/

func dailyTemperatures(T []int) []int {
	n := len(T)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		tem := T[i]
		for len(stack) > 0 && tem > T[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return ans
}
