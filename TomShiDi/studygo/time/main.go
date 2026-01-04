package main

import (
	"fmt"
	"time"
)

func main() {
	ticker()
}

func base() {
	t := time.Now()
	fmt.Println(t)

	// 格式化输出时间，03表示12小时制的小时，15表示24小时制的小时
	format := t.Format("2006-01-02 03:04:05")
	fmt.Println("格式化后= ", format)

	// 获取时间戳
	fmt.Println("时间戳=", t.Unix())
	fmt.Println("纳秒时间戳=", t.UnixNano())

	// 时间戳转时间格式
	t2 := time.Unix(1672531199, 0)
	fmt.Println("时间戳转换时间=", t2.Format("2006-01-02 03:04:05"))

	var str = "2026-01-04 10:32:01"
	var template = "2006-01-02 15:04:05"

	location, _ := time.ParseInLocation(template, str, t.Location())
	fmt.Println("字符串转时间类型= ", location)

	// 时间增减
	fmt.Println("当前时间=", t, " 一小时后时间=", t.Add(time.Hour))
}

// ticker 定时器
func ticker() {
	ticker := time.NewTicker(time.Second)

	n := 5
	for t := range ticker.C {
		fmt.Println("循环批次: ", n, " 当前时间: ", t)
		n--
		if n <= 0 {
			ticker.Stop()
			break
		}
	}
}
