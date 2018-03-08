package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	deferCall()
	parseStudent()
	exercise1()
	exercise2()
	//exercise3()
	exercise4()
	exercise5()
}

func deferCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Printf("[%v]", stus)
}

func exercise1() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func exercise2() {
	t := Teacher{}
	t.ShowA()
}

func exercise3() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func exercise4() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // 参数中的函数会先执行
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

type ThreadSafeSet struct {
	sync.RWMutex
	s []int
}

// 内部迭代出现阻塞。默认初始化时无缓冲区，需要等待接收者读取后才能继续写入。
func (set *ThreadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
			fmt.Print("get:", elem, ",")
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func exercise5() {
	//read()
	unRead()
}

func read() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	closed := false
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Print("read:", v, ",")
			} else {
				closed = true
			}
		}
		if closed {
			fmt.Print("closed")
			break
		}
	}
	fmt.Print("Done")
}

func unRead() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	_ = ch
	time.Sleep(5 * time.Second)
	fmt.Print("Done")
}

