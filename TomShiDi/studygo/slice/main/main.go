package main

import "fmt"

func main() {
	// size决定初始化时实际分配空间的数量，capacity表示扩容阈值
	var s1 []int = make([]int, 2, 4)
	fmt.Printf("s1=%#v len=%d capacity=%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s1=nil?%v \n", s1 == nil)

	var s2 []int
	fmt.Printf("s2=%#v \n", s2)
	fmt.Printf("s2=nil?%v \n", s2 == nil)

	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 切一个空元素的切片
	s3 := arr1[:]
	// 从数组得到的切片，cap等于[切片开始切的位置到数组最后的元素个数]
	fmt.Printf("s3=%#v len=%d capacity=%d \n", s3, len(s3), cap(s3))

	s4 := arr1[3:]
	fmt.Printf("s4=%#v len=%d capacity=%d \n", s4, len(s4), cap(s4))

	// 切片再切
	s5 := s4[3:]
	fmt.Printf("s5=%#v len=%d capacity=%d \n", s5, len(s5), cap(s5))

	// 切片是引用类型，指向底层数组
	// 底层数组修改后，从它上面切出来的切片也会受到影响
	arr1[9] = 888
	fmt.Printf("底层数组修改后s5=%v \n", s5)

	// 修改切片，也会影响源数组对应下边的值
	s5[len(s5)-1] = 0
	fmt.Printf("修改切片后，源数组arr1=%v \n", arr1)
}
