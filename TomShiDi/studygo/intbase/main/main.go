package main

import "fmt"

func main() {
	var intD int = 16
	//var intO int = 020
	//var intX int = 0x10

	fmt.Printf("十进制:%d,八进制：%o,十六进制；%x \n", intD, intD, intD)
	fmt.Printf("变量类型：%T", int16(intD))
}
