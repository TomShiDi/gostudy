package main

import "fmt"

func main() {
	var a int = 10
	var ap *int = &a
	fmt.Printf("a的值=%d，a的地址=%p\n", a, ap)

	b := *ap
	fmt.Printf("通过指针ap获取的值=%d\n", b)
	*ap = 20
	fmt.Printf("修改指针ap指向的值后，a的值=%d\n", a)

	// 创建指针变量
	p := new(int)
	fmt.Printf("指针p的地址=%p，p指向的值=%d\n", p, *p)
	*p = 30
	fmt.Printf("修改指针p指向的值后，p指向的值=%d\n", *p)
}

// 不使用指针，无法修改外部变量的值
func f1(x int) {
	x = 20
}

// 使用指针修改外部变量的值
func f2(x *int) {
	*x = 30
}
