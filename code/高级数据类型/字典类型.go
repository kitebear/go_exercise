package main

import (
	"fmt"
	"sync"
)

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

	// demo 1
	// cMap := make(map[string]int)
	// cMap["北京"] = 1

	// // 指定初始容量
	// cMap = make(map[string]int, 100)
	// cMap["北京"] = 1

	// map 基本操作
	// cMap := map[string]int{}
	// cMap["北京"] = 1 //写
	// code := cMap["北京"] // 读
	// fmt.Println(code)

	// code = cMap["广州"] // 读不存在 key
	// fmt.Println(code)

	// code, ok := cMap["广州"] // 检查 key 是否存在
	// if ok {
	// 	fmt.Println(code)
	// } else {
	// 	fmt.Println("key not exist")
	// }

	// delete(cMap, "北京") // 删除 key
	// fmt.Println(cMap["北京"])

	// 循环和无序性
	// cMap := map[string]int{"北京": 1, "上海": 2, "广州": 3, "深圳": 4}

	// for city, code := range cMap {
	// 	fmt.Printf("%s:%d", city, code)
	// 	fmt.Println()
	// }

	cMap := make(map[string]int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		cMap["北京"] = 1
		wg.Done()
	}()

	go func() {
		cMap["上海"] = 2
		wg.Done()
	}()

	wg.Wait()

	provinces := make(map[string]map[string]int)

	provinces["北京"] = map[string]int{
		"东城区": 1,
		"西城区": 2,
		"朝阳区": 3,
		"海淀区": 4,
	}

	fmt.Println(provinces["北京"])
}
