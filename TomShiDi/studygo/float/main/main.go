package main

import "fmt"

func main() {
	var f1 = 3.1415

	fmt.Printf("默认浮点数类型%T \n", f1)

	f2 := float32(f1)

	fmt.Printf("强制转换成32位浮点型：%T \n", f2)

	var f3 float64 = 1129.6
	var f4 float64 = 1128

	fmt.Println("golang中浮点运算可能存在精度问题：", f3-f4)

}
