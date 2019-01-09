package main

// 第一种用法 起别名
/**
import (
	demo "fmt"
)

func main() {
	demo.Println("hellowork")
}
**/

// 第二种用法 import . 用法
// import (
// 	. "fmt"
// )

// func main() {
// 	Print("helloword")
// }

// 第三种用法 import _ 用法  会执行指定包中的常量和变量 然后执行init方法
import (
	_ "go_exercise/demo"
)

func main() {
}
