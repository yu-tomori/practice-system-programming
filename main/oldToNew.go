package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := bytes.NewBuffer([]byte("aaa"))
	buffer.Write([]byte("\nbbb"))
	io.Copy(buffer, file)
	file2, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	n, err := io.Copy(file2, buffer)
	if err != nil {
		fmt.Printf("\nerr: %v\n", err)
	}
	fmt.Printf("written bytes: %d", n)
}
