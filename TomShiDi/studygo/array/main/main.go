package main

import "fmt"

/**
数组长度是数组类型的一部分
*/
func main() {
	var a1 [3]int = [...]int{1, 2, 3}
	var a2 [4]int = [4]int{1, 2, 3, 4}
	// 顺序初始化，后面自动补默认值
	var a3 [5]int = [5]int{1, 2}
	// 初始化指定下标的元素
	var a4 [5]int = [5]int{0: 4, 4: 1}
	var bollArr [3]bool
	fmt.Printf("a1的类型是%T,a2的类型是%T \n", a1, a2)

	fmt.Printf("a3的值是：%#v \n", a3)
	fmt.Printf("a4的值是：%#v \n", a4)
	// 数组不初始化，bool数组默认false，数字数组默认 0
	fmt.Printf("%#v \n", bollArr)

	//数组遍历
	println("**************")
	for i := 0; i < len(a4); i++ {
		fmt.Printf("a4[%d]=%d \n", i, a4[i])
	}
	println("**************")

	println("**************")
	for index, item := range a4 {
		fmt.Printf("a4[%d]=%d \n", index, item)
	}
	println("**************")

	// 多维数组
	var aa1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var aa2 [2][2]int = [2][2]int{[2]int{1, 2}, [2]int{3, 4}}
	fmt.Printf("aa1的值是：%#v \n", aa1)
	fmt.Printf("aa2的值是：%#v \n", aa2)

	// 多维数组遍历
	println("**************")
	for i := 0; i < len(aa1); i++ {
		for j := 0; j < len(aa1[i]); j++ {
			fmt.Printf("%d \n", aa1[i][j])
		}
	}
	println("**************")

	// 数组是值类型
	b1 := [...]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Printf("b1=%v,b2=%v \n", b1, b2)

	fmt.Printf("b1与b2相等？%v \n", b1 == b2)
}
