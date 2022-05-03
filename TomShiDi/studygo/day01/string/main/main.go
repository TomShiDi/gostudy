package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 = `
		你好
		世界
	`
	println(s1)
	var s2 = "Hello"
	var s3 = "World"
	ss1 := fmt.Sprintf("%s %s", s2, s3)
	fmt.Printf("fmt.Sprintf的值是： %s \n", ss1)

	split := strings.Split(ss1, " ")
	fmt.Printf("Split值是：%#v \n", split)

	fmt.Printf("ss1是否包含Hello: %#v \n", strings.Contains(ss1, s2))
	fmt.Printf("ss1前缀是否为Hello: %#v \n", strings.HasPrefix(ss1, s2))
	fmt.Printf("ss1后缀是否为World: %#v \n", strings.HasSuffix(ss1, s3))
	fmt.Printf("World在ss1中的位置是：%#v \n", strings.Index(ss1, s3))
	fmt.Printf("ss1中o最后出现的位置: %#v \n", strings.LastIndex(ss1, "o"))
	fmt.Printf("Join拼接后的字符: %#v \n", strings.Join(split, "+"))

	ss1 = ss1 + " 中文"
	for _, v := range ss1 {
		fmt.Printf("%c\n", v)
	}
}
