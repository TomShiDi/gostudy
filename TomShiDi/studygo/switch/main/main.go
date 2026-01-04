package main

import "fmt"

func main() {
	var enum int = 5
	switch enum {
	case 1:
		fmt.Println("类型一")
		break
	case 2:
		fmt.Println("类型二")
		break
	case 3:
		fmt.Println("类型三")
		break
	case 4:
		fmt.Println("类型四")
		break
	default:
		fmt.Println("其他选项")
	}

	fmt.Println("---------------")
	switch enum := 1; enum {
	case 1:
		fmt.Println("类型一")
		break
	case 2:
		fmt.Println("类型二")
		break
	case 3:
		fmt.Println("类型三")
		break
	case 4:
		fmt.Println("类型四")
		break
	default:
		fmt.Println("其他选项")
	}

	fmt.Println("---------------")
	switch enum := 7; enum {
	case 1, 2, 3, 4, 5:
		fmt.Println("有效选项")
		break
	default:
		fmt.Println("无效选项")
	}

	fmt.Println("---------------")
	fmt.Scanf("%d", &enum)
	switch {
	case enum > 100:
		fmt.Println("大于1")
		fallthrough
	case enum > 2:
		fmt.Println("大于2")
	default:
		fmt.Println("其他值")
	}
}
