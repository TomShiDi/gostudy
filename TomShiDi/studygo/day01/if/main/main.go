package main

func main() {
	//var age int = 19
	//if age > 18 {
	//	fmt.Println("已成年")
	//} else if age > 50 {
	//	println("中年了")
	//} else if age > 200 {
	//	println("升天了")
	//} else {
	//	println("未成年")
	//}

	// 此处age只在if作用域内生效
	if age := 19; age > 19 {
		println("成年了")
	} else {
		println("未成年")
	}
}
