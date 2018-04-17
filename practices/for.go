package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan struct{})

	for _, v := range s {
		go func() {
			fmt.Println(v)  // 错误
			ch <- struct{}{}
		}()
	}

	// 等待所有的协程结束
	for range s {
		<- ch
	}

	/*for _, v := range s {
		go func(v int) {
			fmt.Println(v)  // ok, v是闭包的形参，值与循环变量一致
			ch <- struct{}{}
		}(v)  // 将循环变量作为实参传入
	}*/
}