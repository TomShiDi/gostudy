package main

import (
	"fmt"
	"strconv"
)

func main() {
	complex()
}

func base() {
	map1 := make(map[string]int, 8)
	map1["Tom"] = 18
	map1["Shi"] = 19
	map1["Di"] = 20

	// 遍历map
	for key, value := range map1 {
		fmt.Println("key=", key, " value=", value)
	}

	map2 := map[string]string{
		"name": "tomshidi",
		"age":  "18",
	}
	fmt.Println("map2=", map2)
	fmt.Println("map2[name]=", map2["name"])

	// 删除元素
	delete(map2, "age")
	fmt.Println("删除元素后map2=", map2)
	v, ok := map2["age"]
	fmt.Printf("Map2删除age后进行查询，age=%#v,ok=%v", v, ok)
}

func complex() bool {
	userInfos := make([]map[string]string, 5)
	for i := 0; i < len(userInfos); i++ {
		v := make(map[string]string)
		userInfos[i] = v
		v["name"] = "TomShiDi_" + strconv.Itoa(i)
		v["age"] = strconv.Itoa(100 + i)
	}
	fmt.Printf("userInfos=%v \n", userInfos)

	return true
}
