package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(name string) {
	data, err := ioutil.ReadFile(name)
	errCheck(err)
	fmt.Println("data: ", string(data))
}

func main() {
	path := "writer.txt"

	// Block 1
	// newFile, err := os.Create(path)
	// if err != nil {
	// 	fmt.Println("error is :", err)
	// 	return
	// }

	// defer newFile.Close()

	// info, err := os.Stat(path)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Println("the file is not exist")
	// 	} else {
	// 		fmt.Print("error is: ", err)
	// 	}
	// 	return
	// }

	// fmt.Println("the file name: ", info.Name())

	// Block 2
	// content := "Hello World"

	// err := ioutil.WriteFile(path, []byte(content), 0644)
	// if err != nil {
	// 	fmt.Println("ioutil.WriteFile err is: ", err)
	// 	return
	// }

	// data, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	fmt.Println("ioutil.ReadFile(%v)", path, err)
	// 	return
	// }

	// fmt.Println("File content is: ", string(data))

	// Block 3
	// content := "hello"
	// newContent := "world"
	// newFile, err := os.Create(path)
	// errCheck(err)

	// n1, err := newFile.WriteAt([]byte(content), 0)
	// errCheck(err)
	// fmt.Println("n1 : ", n1)
	// readFile(path)

	// n2, err := newFile.WriteAt([]byte(newContent), 6)
	// errCheck(err)
	// fmt.Println("n2 : ", n2)
	// readFile(path)

	content := "hello"
	newFile, err := os.Create(path)
	errCheck(err)

	bufferWriter := bufio.NewWriter(newFile)

	for _, v := range content {
		n, err := bufferWriter.WriteString(string(v))
		errCheck(err)
		fmt.Println("n is : ", n)
	}

	readFile(path) // empty

	unflushSize := bufferWriter.Buffered()
	fmt.Println("unflushSize: ", unflushSize)

	bufferWriter.Flush()

	readFile(path)
}
