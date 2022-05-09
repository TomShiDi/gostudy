package main

import "fmt"

func main() {
	var f1 = 3.1415

	fmt.Printf("默认浮点数类型%T \n", f1)

	f2 := float32(f1)

	fmt.Printf("强制转换成32位浮点型：%T", f2)

}
