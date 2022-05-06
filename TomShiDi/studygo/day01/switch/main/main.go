package main

func main() {
	var enum int = 5
	switch enum {
	case 1:
		println("类型一")
		break
	case 2:
		println("类型二")
		break
	case 3:
		println("类型三")
		break
	case 4:
		println("类型四")
		break
	default:
		println("其他选项")
	}

	println("---------------")
	switch enum := 1; enum {
	case 1:
		println("类型一")
		break
	case 2:
		println("类型二")
		break
	case 3:
		println("类型三")
		break
	case 4:
		println("类型四")
		break
	default:
		println("其他选项")
	}

	println("---------------")
	switch enum := 7; enum {
	case 1, 2, 3, 4, 5:
		println("有效选项")
		break
	default:
		println("无效选项")
	}

	println("---------------")
	switch {
	case enum > 1:
		println("大于1")
		fallthrough
	case enum > 2:
		println("大于2")
	default:
		println("其他值")
	}
}
