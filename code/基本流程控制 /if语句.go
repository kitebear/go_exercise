package main

import "fmt"

func main() {
	var number int = 5
	if number += 4; 10 > number {
		number *= 4
		number += 3
		fmt.Println(number)
	} else if 10 < number {
		number -= 2
		fmt.Print(number)
	}
	fmt.Println(number)
}
