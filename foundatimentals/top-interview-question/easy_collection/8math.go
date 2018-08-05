package leetcode

import (
	"math"
	"strconv"
	"strings"
)

// Fizz Buzz
/*
	写一个程序，输出从 1 到 n 数字的字符串表示。
	1. 如果 n 是3的倍数，输出“Fizz”；
	2. 如果 n 是5的倍数，输出“Buzz”；
	3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
	示例：
	n = 15,
	返回:
	[
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz"
	]
*/

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i, fizz, buzz := 1, 0, 0; i <= n; i++ {
		fizz++
		buzz++
		if fizz == 3 && buzz == 5 {
			res = append(res, "FizzBuzz")
			fizz, buzz = 0, 0
		} else if fizz == 3 {
			res = append(res, "Fizz")
			fizz = 0
		} else if buzz == 5 {
			res = append(res, "Buzz")
			buzz = 0
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

func fizzBuzz2(n int) []string {
	rv := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		j, k := i%3, i%5
		if j == 0 && k == 0 {
			rv = append(rv, "FizzBuzz")
		} else if j == 0 {
			rv = append(rv, "Fizz")
		} else if k == 0 {
			rv = append(rv, "Buzz")
		} else {
			rv = append(rv, strconv.Itoa(i))
		}
	}
	return rv
}

// Count Primes 计数质数
/*
	统计所有小于非负整数 n 的质数的数量。
	示例:
	输入: 10
	输出: 4
	解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/
func countPrimes(n int) int {
	notPrime := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}
	return count
}

func countPrimes2(n int) int {
	var sum int
	f := make([]bool, n)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if f[i] == false {
			for j := i * i; j < n; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if f[i] == false {
			sum++
		}
	}
	return sum
}

// Power of Three
/*
	给定一个整数，写一个函数来判断它是否是 3 的幂次方。
	示例 1:
		输入: 27
		输出: true
		示例 2:
		输入: 0
		输出: false
		示例 3:
		输入: 9
		输出: true
		示例 4:
		输入: 45
		输出: false
	进阶：
		你能不使用循环或者递归来完成本题吗？
*/

// 1162261467 is 3^19,  3^20 is bigger than int
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func isPowerOfThree2(n int) bool {
	return n > 0 && (n == 1 || (n%3 == 0 && isPowerOfThree2(n/3)))
}

func isPowerOfThree3(n int) bool {
	if n > 1 {
		for n%3 == 0 {
			n /= 3
		}
	}

	return n == 1
}

// Roman to Integer 罗马数字转整数
/*
	罗马数字包含以下七种字符：I， V， X， L，C，D 和 M。

	字符          数值
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
	例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

	通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

	I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

	示例 1:

	输入: "III"
	输出: 3
	示例 2:

	输入: "IV"
	输出: 4
	示例 3:

	输入: "IX"
	输出: 9
	示例 4:

	输入: "LVIII"
	输出: 58
	解释: C = 100, L = 50, XXX = 30, III = 3.
	示例 5:

	输入: "MCMXCIV"
	输出: 1994
	解释: M = 1000, CM = 900, XC = 90, IV = 4.
*/
func romanToInt(s string) int {
	result := 0

	if strings.Index(s, "IV") != -1 {
		result -= 2
	}
	if strings.Index(s, "IX") != -1 {
		result -= 2
	}
	if strings.Index(s, "XL") != -1 {
		result -= 20
	}
	if strings.Index(s, "XC") != -1 {
		result -= 20
	}
	if strings.Index(s, "CD") != -1 {
		result -= 200
	}
	if strings.Index(s, "CM") != -1 {
		result -= 200
	}

	for _, v := range s {
		if v == 'I' {
			result += 1
		}
		if v == 'V' {
			result += 5
		}
		if v == 'X' {
			result += 10
		}
		if v == 'L' {
			result += 50
		}
		if v == 'C' {
			result += 100
		}
		if v == 'D' {
			result += 500
		}
		if v == 'M' {
			result += 1000
		}
	}
	return result
}

func romanToInt2(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var out, prev int
	for i := len(s) - 1; i >= 0; i-- {
		cur := values[s[i]]
		if cur < prev {
			out -= cur
		} else {
			out += cur
		}
		prev = cur
	}
	return out
}
