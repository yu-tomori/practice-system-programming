package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	_, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("%s is not exist", os.Args[1])
	} else if err != nil {
		fmt.Printf("err type: %T", err)
		panic(err)
	} else {
		fmt.Printf("%s is exist", os.Args[1])
	}
}
