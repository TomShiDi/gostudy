package main

import "fmt"

func main() {
	sliceMethod()
}

func base() {
	// size决定初始化时实际分配空间的数量，capacity表示扩容阈值
	var s1 []int = make([]int, 2, 4)
	fmt.Printf("s1=%#v len=%d capacity=%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s1=nil?%v \n", s1 == nil)

	// 切一个空元素的切片
	// 一个nil值的切片长度和容量都是0，但是长度和容量都是0的切片不一定是nil
	var s2 []int
	fmt.Printf("s2=%#v \n", s2)
	fmt.Printf("s2=nil?%v \n", s2 == nil)

	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s3 := arr1[:]
	// 从数组得到的切片，cap等于[切片开始切的位置到数组最后的元素个数]
	fmt.Printf("s3=%#v len=%d capacity=%d \n", s3, len(s3), cap(s3))

	s4 := arr1[3:]
	fmt.Printf("s4=%#v len=%d capacity=%d \n", s4, len(s4), cap(s4))

	// 切片再切
	s5 := s4[3:4]
	fmt.Printf("s5=%#v len=%d capacity=%d \n", s5, len(s5), cap(s5))

	// 切片是引用类型，指向底层数组
	// 底层数组修改后，从它上面切出来的切片也会受到影响
	arr1[9] = 888
	fmt.Printf("底层数组修改后s5=%v \n", s5)

	// 修改切片，也会影响源数组对应下标的值
	s5[len(s5)-1] = 0
	fmt.Printf("修改切片后，源数组arr1=%v \n", arr1)

	// make函数构建切片
	// 切片是引用类型，数据保存在底层数组中
	ms := make([]int, 0, 2)
	fmt.Printf("ms=%#v,ms的len=%d,cap=%d \n", ms, len(ms), cap(ms))

	// 切片的类型赋值是浅拷贝
	s6 := []int{1, 2, 3}
	s7 := s6
	s6[0] = 1000

	fmt.Printf("s6=%v,s7=%v \n", s6, s7)
}

func sliceMethod() {
	sliceA := make([]int, 5, 10)
	fmt.Printf("sliceA=%v, len=%d, cap=%d \n", sliceA, len(sliceA), cap(sliceA))

	// 使用append函数向切片添加元素
	sliceA = append(sliceA, 1, 2, 3)
	fmt.Printf("sliceA=%v, len=%d, cap=%d \n", sliceA, len(sliceA), cap(sliceA))

	// 使用append函数添加另一个切片，需要展开
	sliceB := make([]int, 0, 2)
	sliceB = append(sliceB, sliceA...)
	fmt.Printf("sliceB=%v, len=%d, cap=%d \n", sliceB, len(sliceB), cap(sliceB))

	// 切片的拷贝
	sliceC := make([]int, 20)
	n := copy(sliceC, sliceB)
	fmt.Printf("sliceC=%v, len=%d, cap=%d, n=%d \n", sliceC, len(sliceC), cap(sliceC), n)

	// 删除切片元素 4
	sliceD := []int{1, 2, 3, 4, 5, 6}
	sliceD = append(sliceD[:3], sliceD[4:]...)
	fmt.Printf("sliceD=%v \n", sliceD)
}
