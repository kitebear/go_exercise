package main

import (
	"fmt"
)

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for _, v := range []int{1, 2, 3, 4} {
		fmt.Printf("%d: %s \n", v, map1[v])
	}

	numbers := []int{1, 2, 3}

	for i, v := range numbers {
		fmt.Printf("Numbers: i %d, v %d \n", i, v)
	}

	cityCodes := map[string]int{
		"北京": 1,
		"上海": 2,
	}

	for i, v := range cityCodes {
		fmt.Printf("%s is %d\n", i, v)
	}

	for i, v := range numbers {
		if v == 4 {
			break
		}

		if v%2 == 0 {
			continue
		}

		fmt.Printf("numbers[%d] is %d\n", i, v)
	}
}
