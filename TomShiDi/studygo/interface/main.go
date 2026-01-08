package main

import "fmt"

// 空接口，可以接收任意类型的值
type Any interface{}

// 结构体的中的方法是值类型接收者，那么实例化后结构体值类型和结构体指针类型的变量都可以调用该方法
type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "关机")
}

// 结构体的中的方法是指针类型接收者，那么实例化后结构体指针类型的变量才能调用该方法
//func (p *Phone) Start() {
//	fmt.Println(p.Name, "启动")
//}
//
//func (p *Phone) Stop() {
//	fmt.Println(p.Name, "关机")
//}

// show 使用空接口作为参数
func show(a interface{}) {
	fmt.Printf("a=%v, a类型=%T\n", a, a)
}

func main() {
	var p Usber = Phone{"iPhone"}
	p.Start()

	str := "hello"
	var a Any = str
	fmt.Printf("a=%v, a类型=%T\n", a, a)

	m := map[string]interface{}{"name": "Tom", "age": 30, "isStudent": false}
	show(m)

	// 类型断言
	var any interface{}
	any = "字符串"
	v, ok := any.(string)
	if ok {
		fmt.Println("是字符串类型，v=", v)
	} else {
		fmt.Println("不是字符串类型")
	}

	// 类型判断
	switch any.(type) {
	case string:
		fmt.Println("字符串类型")
	case int:
		fmt.Println("整数类型")
	default:
		fmt.Println("其他类型")
	}

	//
	var pp = &Phone{"Samsung"}
	pp.Start()
}
