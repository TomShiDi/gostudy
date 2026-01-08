package main

import "fmt"

type Animal interface {
	SetName(string)
	GetName() string
}

type Behavior interface {
	Say()
}

// 接口嵌套
type Nesting interface {
	Animal
	Behavior
}

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d Dog) GetName() string {
	return d.name
}

func (d Dog) Say() {
	fmt.Println(d.GetName(), " 汪汪汪")
}

func main() {
	var d1 Animal = &Dog{"阿巴阿巴"}
	fmt.Println(d1.GetName())
	var d2 Behavior = Dog{"小黑"}
	d2.Say()
}
