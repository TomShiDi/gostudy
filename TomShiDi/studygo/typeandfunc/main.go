package main

import "fmt"

type calc func(int, int) int

type myInt int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// calcFunc 使用函数类型作为参数
func calcFunc1(a, b int, c func(a, b int) int) int {
	return c(a, b)
}

func calcFunc2(a, b int, c calc) int {
	return c(a, b)
}

// returnFunc 使用函数类型作为返回值
func returnFunc(o string) func(int, int) int {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x int, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

// 闭包使得方法中的i，对于某个闭包函数实例来说是全局的
func closure() func(int) int {
	i := 10
	return func(x int) int {
		i = i + x
		return i
	}
}

func main() {
	c := add
	fmt.Println("1+2=", c(1, 2))

	// 使用自定义类型
	var a myInt = 10
	fmt.Printf("a类型是=%T\n", a)

	// 使用函数类型作为参数
	result := calcFunc1(10, 5, sub)
	fmt.Println("10-5=", result)

	// 使用函数类型作为参数，传入匿名函数
	result = calcFunc2(20, 10, func(x int, y int) int {
		return x * y
	})
	fmt.Println("20*10=", result)

	// 使用函数类型作为返回值
	f := returnFunc("+")
	if f != nil {
		fmt.Println("5+3=", f(5, 3))
	}

	func(x, y int) {
		fmt.Println("匿名自执行函数,x=", x, "y=", y)
	}(1, 2)

	fn := closure()
	fmt.Println("闭包执行第一次+10后= ", fn(10))
	fmt.Println("闭包执行第二次+10后= ", fn(10))
	fmt.Println("闭包执行第三次+10后= ", fn(10))
}
