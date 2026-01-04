package main

import (
	"errors"
	"fmt"
)

// defer 延迟执行
// defer语句会将其后面的函数推迟到包含它的函数返回之后才执行
// 推迟调用的函数其参数会立即求值，但直到上层函数返回前该函数都不会被调用
// 多个defer语句会被压入一个栈中，按后进先出顺序执行
func f1() {
	fmt.Println("开始执行f1")
	defer func() {
		fmt.Println("执行defer闭包自执行函数")
	}()
	fmt.Println("结束执行f1")
}

// 匿名返回值 0
func f2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

// 命名返回值 6
func f3() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	defer fmt.Println("calc函数执行完毕，index=", index, "ret=", ret)
	return ret
}

func complexDefer() {
	x := 1
	y := 2
	defer calc("AAA", x, y) // 这里的x和y的值会被立即求值，传入的是1和2
	x = 10
	defer calc("BBB", x, y) // 这里的x和y的值会被立即求值，传入的是10和2
	y = 20
}

// recover 捕获异常
// recover内置函数可以让处于panic状态中的goroutine恢复过来
// recover只有在被defer调用的函数中才有效
func panicDefer() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic err=", err)
		}
	}()
	panic("抛出异常")
}

func doReadFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件出错")
	}
}

func readFile() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到异常：", err)
		}
	}()
	err := doReadFile("abc.txt")
	if err != nil {
		panic(err)
	}
}

func main() {
	readFile()
}
