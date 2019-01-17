package main

import (
	"fmt"
	"math/rand"
)

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is a int\n", v)
	case string:
		fmt.Printf("%v is a string\n", v)
	default:
		fmt.Printf("%v's type is unkown\n", v)
	}
}

func main() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4)]; interface{}(v).(type) {
	case uint, int32:
		fmt.Printf("Case A.")
	case byte:
		fmt.Printf("Case B.")
	default:
		fmt.Println("Unknown!")
	}

	age := 7
	// switch 的 case 条件可以是多个值
	switch age {
	case 7, 8, 9, 10, 11, 12:
		fmt.Println("It's primary school")
	case 13, 14, 15:
		fmt.Println("It's middle school")
	case 16, 17, 18:
		fmt.Println("It's high school")
	default:
		fmt.Println("The age is unkown")
	}

	//还可以使用 if..else 作为 case 条件，例如：
	switch {
	case age >= 6 && age <= 12:
		fmt.Println("It's primary school")
	case age >= 13 && age <= 15:
		fmt.Println("It's middle school")
	case age >= 16 && age <= 18:
		fmt.Println("It's high school")
	default:
		fmt.Println("The age is unkown")
	}

	//  使用 switch 对 interface{} 进行断言
	checkType(8)
	checkType("hello, world")
}
