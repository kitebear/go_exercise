package main

import "fmt"

func main() {
	mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
	mm2["scala"] = 25
	mm2["erlang"] = 50
	mm2["python"] = 0
	fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["python"])

	// var cMap map[string]int // 只定义, 此时 cMap 为 nil
	// fmt.Println(cMap == nil)
	// cMap["北京"] = 1 // 报错，因为 cMap 为 nil

	// var cMap map[string]int

	cMap := make(map[string]int)
	cMap["北京"] = 1

	// 指定初始容量
	cMap = make(map[string]int, 100)
	cMap["北京"] = 1
}
