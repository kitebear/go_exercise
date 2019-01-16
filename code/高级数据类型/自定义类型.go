package main

import "fmt"

// 单次声明
type City string
type Age int

func printAge(age int) {
	fmt.Println("Age is", age)
}

func main() {
	// 批量声明
	type (
		B0 = int8
		B1 = int16
		B2 = int32
		B3 = int64
	)

	type (
		A0 int8
		A1 int16
		A2 int32
		A3 int64
	)

	// 简单示例
	city := City("上海")
	fmt.Println(city)
	fmt.Println("I live in", city+" 上海") //  字符串拼接
	fmt.Println(len(city))               // len 方法
	middle := Age(12)

	if middle >= 12 {
		fmt.Println("Middle is bigger than 12")
	}

	// 强制转换类型
	printAge(int(middle))
}
