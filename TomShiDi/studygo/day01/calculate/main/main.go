package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	fmt.Printf("a=%d,b=%d\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%d,b=%d\n", a, b)
	// a++是单独的语句，不能与其他运算一起使用，例如不能用：a = a++

	// 逻辑运算符：&& || !
	if a > 0 && a < 3 {
		println("三岁内")
	}

	// 位运算 &(与) |(或) ^(异或) << >>
	fmt.Printf("二进制单位跨度：%d", 1<<10)
}
