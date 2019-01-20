package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Student struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Results []Result `json:"results"`
}

type Result struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

func main() {
	data, _ := ioutil.ReadFile("data.json")
	var s Student
	// 反序列化
	err := json.Unmarshal(data, &s)

	fmt.Println("data error is : ", err)
	fmt.Printf("Unmarshal %v\n", s)

	data2, err := json.Marshal(s)
	fmt.Println("data2 error is : ", err)

	ioutil.WriteFile("data2.json", data2, 0755)
}
