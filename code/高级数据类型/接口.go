package main

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name    string
	Age     int8
	Address string
}

func (c *Cat) Grow() {
	c.Age++
}

func (c *Cat) Move(str string) string {
	return str
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)
}
