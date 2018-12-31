package main

import (
	"errors"
	"fmt"
)

func innerFunc() {
	fmt.Println("Enter innerFunc")
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Fatal error: %s\n", p)
		}
	}()
	panic(errors.New("Occur a panic!"))
	fmt.Println("Quit innerFunc")
}

func outerFunc() {
	fmt.Println("Enter outerFunc")
	innerFunc()
	fmt.Println("Quit outerFunc")
}

func main() {
	fmt.Println("Enter main")
	outerFunc()
	fmt.Println("Quit main")
}
