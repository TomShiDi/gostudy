package main

import "fmt"

var (
	name    string
	age     int
	isOK    bool
	arr     []int = make([]int, 1)
	arrAuto       = [...]int{1, 2, 3}
)

func main() {
	name = "TomShiDi"
	age = 18
	isOK = true

	fmt.Printf("name=%s,age=%d,isOK=%t", name, age, isOK)
}
