package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Block 1
	// var (
	// 	a = "1"
	// 	b int
	// )

	// err := json.Unmarshal([]byte(a), &b)
	// fmt.Println("Unmarshal err is :", err)
	// fmt.Printf("Unmarshal result is %T %d \n", b, b)

	// data, err := json.Marshal(b)
	// fmt.Println("Marshal err is", err)
	// fmt.Printf("Marshal result is %s \n", string(data))

	var (
		d1 = `true`
		v1 bool
	)

	json.Unmarshal([]byte(d1), &v1)
	printHelper("d1", v1)

	var (
		d2 = `"hello,word"`
		v2 string
	)
	json.Unmarshal([]byte(d2), &v2)
	printHelper("d2", v2)

	var (
		d3 = `1`
		v3 int
	)

	json.Unmarshal([]byte(d3), &v3)
	printHelper("d3", v3)

	var (
		d4 = `3.14`
		v4 float64
	)

	json.Unmarshal([]byte(d4), &d4)
	printHelper("d4", v4)

	var (
		d5 = `[1,2,3]`
		v5 []int
	)

	json.Unmarshal([]byte(d5), &v5)
	printHelper("d5", v5)

	var (
		d6 = `{"a": "b"}`
		v6 map[string]string
	)

	json.Unmarshal([]byte(d6), &v6)
	printHelper("d6", v6)
}

func printHelper(name string, v interface{}) {
	fmt.Printf("%s unmarshal value is %T, %v\n", name, v, v)
}
