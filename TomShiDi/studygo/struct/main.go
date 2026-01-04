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

type Remark struct {
	note string
}

type ComplexPerson struct {
	p      people
	Remark                   // 匿名字段，使用时可以直接通过ComplexPerson访问people和Remark的字段和方法
	hobby  []string          // 结构体中切片类型初始值为nil，需要初始化才能使用
	others map[string]string // 结构体中map类型初始值为nil，需要初始化才能使用
}

// 父结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Println(a.Name, " 在跑")
}

type Dog struct {
	Age int
	Animal
}

func (d Dog) say() {
	fmt.Println(d.Name, " 在汪汪汪")
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

	cp := ComplexPerson{}
	fmt.Printf("cp=%#v\n", cp)
	cp.p.name = "tomshidi"
	cp.p.age = 18
	cp.hobby = []string{"吃饭", "睡觉", "打豆豆"}
	cp.others = make(map[string]string)
	cp.others["address"] = "广东"
	// 访问匿名字段Remark的note字段，首先会在ComplexPerson中查找note字段，找不到再去Remark中查找
	cp.note = "这是一个匿名字段"
	fmt.Printf("cp=%#v\n", cp)

	// 结构体嵌套、继承
	dog := Dog{
		Age: 3,
		Animal: Animal{
			Name: "旺财",
		},
	}
	dog.run()
	dog.say()
}
