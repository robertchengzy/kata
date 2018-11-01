package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stop3()
}

func stop1() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了。。。")
				return
			default:
				fmt.Println("goroutine监控中。。。")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("通知监控停止了。。。")
	stop <- true

	time.Sleep(5 * time.Second)
}

func stop2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了。。。")
				return
			default:
				fmt.Println("goroutine监控中。。。")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("通知监控停止了。。。")
	cancel()

	time.Sleep(5 * time.Second)
}

func stop3() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "监控1")
	go watch(ctx, "监控2")
	go watch(ctx, "监控3")

	time.Sleep(10 * time.Second)
	fmt.Println("通知监控停止了。。。")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了。。。")
			return
		default:
			fmt.Println(name, "goroutine监控中。。。")
			time.Sleep(2 * time.Second)
		}
	}
}
