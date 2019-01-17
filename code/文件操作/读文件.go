package main

import (
	"fmt"
	"os"
)

func main() {
	// Block 1
	// dat, err := ioutil.ReadFile("./temp.go")
	// fmt.Println(err)
	// fmt.Println(string(dat))

	f, _ := os.Open("./temp.go")
	defer f.Close()

	// Block 2
	// buf := make([]byte, 16)
	// f.Read(buf)

	// fmt.Println(string(buf))

	b1 := make([]byte, 2)

	f.Seek(5, 0)
	f.Read(b1)
	fmt.Println(string(b1))
}
