package main

import (
	"bufio"
	"fmt"
	"io"
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

	// Block 3
	// b1 := make([]byte, 2)
	// f.Seek(5, 0)
	// f.Read(b1)
	// fmt.Println(string(b1))

	// Block 4
	// b1 := make([]byte, 2)
	// f.ReadAt(b1, 5)
	// fmt.Println(string(b1))

	// Block 5
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		buff := make([]byte, 2)
	// 		// f.Seek(5, 0)
	// 		// f.Read(buff)
	// 		f.ReadAt(buff, 5)
	// 		fmt.Println(string(buff))
	// 	}()
	// }
	// time.Sleep(2 * time.Second)

	// Block 6
	// buff := make([]byte, 16)

	// for {
	// 	n, err := f.Read(buff)

	// 	if n == 16 {
	// 		fmt.Print(string(buff))
	// 	} else {
	// 		fmt.Print(string(buff[0:n]))
	// 	}

	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	// Block7
	r := bufio.NewReader(f)
	totalLine := 0

	for {
		_, isPrefix, err := r.ReadLine()

		if !isPrefix {
			totalLine++
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Println("the total line is:", totalLine)
}
