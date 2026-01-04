package main

import (
	"encoding/json"
	"fmt"
)

// 私有属性不能被json包访问
type Student struct {
	ID     int `json:"id"` // 指定序列化后的key名称
	Name   string
	Gender string
	Sno    string
}

type Class struct {
	Title    string
	Students []Student
}

func main() {
	var s1 Student = Student{
		ID:     1,
		Name:   "Tom",
		Gender: "Male",
		Sno:    "1001",
	}
	// 结构体变量转换为json字符串
	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("jsonStr=%#v\n", jsonStr)

	var s2 Student
	err := json.Unmarshal([]byte(`{"id":2,"Name":"Tohka","Gender":"Female","Sno":"1002"}`), &s2)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	fmt.Printf("s2=%#v\n", s2)

	c := Class{
		Title:    "一班",
		Students: []Student{s1, s2},
	}
	cByte, _ := json.Marshal(c)
	fmt.Printf("cByte=%s\n", string(cByte))
}
