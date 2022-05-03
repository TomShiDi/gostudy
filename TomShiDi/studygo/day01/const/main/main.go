package main

import "fmt"

const pi = 3.141592654

const (
	STATUS_SUCCESS  = 200
	STATUS_NOTFOUND = 404
	STATUS_ERROR    = 500
)

const (
	n1 = 100
	n2
	n3
)

/**
iota在const中首次出现时被重置为0，此后每新增一行常量声明iota值会加一
*/
const (
	ENUM1 = iota
	ENUM2
	ENUM3
)

const (
	B1 = iota
	B2 = iota
	_  = iota
	B3 = iota
)

const (
	C1 = iota
	C2 = 100
	C3 = iota
	C4
)

const (
	D1, D2 = iota + 1, iota + 2
	D3, D4 = iota + 1, iota + 2
)

/**
定义数量级
*/
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Printf("ENUM1=%d,ENUM2=%d,ENUM3=%d \n", ENUM1, ENUM2, ENUM3)
	fmt.Printf("B1=%d,B2=%d,B3=%d \n", B1, B2, B3)
	fmt.Printf("C1=%d,C2=%d,C3=%d,C4=%d \n", C1, C2, C3, C4)
	fmt.Printf("D1=%d,D2=%d,D3=%d,D4=%d \n", D1, D2, D3, D4)
	fmt.Printf("KB=%d,MB=%d,GB=%d,TB=%d,PB=%d \n", KB, MB, GB, TB, PB)
}
