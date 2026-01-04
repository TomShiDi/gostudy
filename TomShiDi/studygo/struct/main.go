package main

import "fmt"

// 首字母小写，表示结构体私有
type people struct {
	name string
	age  int
}

// 结构体方法
func (p people) printInfo() {
	fmt.Printf("name=%s, age=%d\n", p.name, p.age)
}

// 结构体指针方法
func (p *people) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	// 直接创建结构体变量
	p1 := people{
		name: "Tom",
		age:  1000,
	}
	// 结构体是值类型，修改p11不会影响p1
	p11 := p1
	p11.name = "Jack"
	fmt.Printf("p1=%#v\n", p1)

	// 使用new关键字创建结构体变量
	p2 := new(people)
	// 等价于 (*p2).name = "tohka"
	p2.name = "tohka"
	p2.age = 18
	fmt.Printf("p2=%#v\n", p2)

	// 使用取地址符号&创建结构体变量
	p3 := &people{}
	p3.name = "yoxino"
	fmt.Printf("p3=%#v\n", p3)

	p4 := &people{
		"tomshidi",
		20,
	}
	//fmt.Printf("p4=%#v\n", p4)
	p4.printInfo()
	p4.setInfo("你好", 18)
	p4.printInfo()
}
