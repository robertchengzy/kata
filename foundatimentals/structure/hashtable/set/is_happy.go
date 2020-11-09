package set

/*
快乐数
编写一个算法来判断一个数 n 是不是快乐数。
「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，
也可能是 无限循环 但始终变不到 1。如果 可以变为 1，那么这个数就是快乐数。
如果 n 是快乐数就返回 True；不是，则返回 False。
示例：
	输入：19
	输出：true
	解释：
	1*1 + 9*9 = 82
	8*8 + 2*2 = 68
	6*6 + 8*8 = 100
	1*1 + 0*0 + 0*0 = 1
*/

func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func isHappy1(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}
