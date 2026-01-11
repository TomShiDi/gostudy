package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("协程1", i, "：你好，Goroutine！")
		time.Sleep(time.Millisecond * 100)
	}
	// 协程的计数器-1
	wg.Done()
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("协程2", i, "：你好，Goroutine！")
		time.Sleep(time.Millisecond * 100)
	}
	// 协程的计数器-1
	wg.Done()
}

func main() {
	// 主协程的计数器+2
	wg.Add(2)
	// 启动一个新的协程执行test函数
	go test1()
	// 启动一个新的协程执行test函数
	go test2()

	for i := 0; i < 10; i++ {
		fmt.Println("主线程", i, "：你好，Goroutine！")
		time.Sleep(time.Millisecond * 100)
	}

	// 等待所有的协程执行完毕
	wg.Wait()
}
