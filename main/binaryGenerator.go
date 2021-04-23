package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	p := make([]byte, 1024)
	n, err := rand.Read(p)
	fmt.Fprintf(os.Stdout, "%d bytes written\n", n)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("binary.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(p)
}
