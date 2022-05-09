package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		// 跳过2的倍数
		if i%2 == 0 {
			continue
		}
		println(i)
	}

	var str string = "Hello世界"

	for i, v := range str {
		fmt.Printf("[%d]=%c\t", i, v)
	}
}
