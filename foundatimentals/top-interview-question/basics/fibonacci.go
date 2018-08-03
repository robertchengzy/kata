package basics

import "fmt"

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibarray(term int) []int {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1

	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}

func fibonaccifunc() func() int {
	back1, back2 := 0, 1
	return func() int {
		// 重新赋值
		back1, back2 = back2, back1+back2
		return back1
	}
}

func fibonaccifuncTest() {
	f := fibonaccifunc()      //  返回一个闭包函数
	for i := 0; i < 10; i++ { // 检测下前10个值
		fmt.Println(f())
	}
}
