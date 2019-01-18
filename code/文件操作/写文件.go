package main

import (
	"fmt"
	"os"
)

func main() {
	path := "writer.txt"

	newFile, err := os.Create(path)
	if err != nil {
		fmt.Println("error is :", err)
		return
	}

	defer newFile.Close()

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("the file is not exist")
		} else {
			fmt.Print("error is: ", err)
		}
		return
	}

	fmt.Println("the file name: ", info.Name())
}
