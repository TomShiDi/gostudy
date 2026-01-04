package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

func main() {
	d := decimal.NewFromFloat(100.114)
	fmt.Printf("decimal=%#v\n", d)

	jsonStr := `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"friends":[{"first":"James","last":"Murphy"},{"first":"Roger","last":"Craig"}]}`
	fmt.Printf("children.1 = %s\n", gjson.Get(jsonStr, "children.0").String())
	fmt.Printf("friends.0.last = %s\n", gjson.Get(jsonStr, "friends.0.last").String())
}
