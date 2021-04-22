package main

import (
	"compress/gzip"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gx")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	defer writer.Close()
	writer.Header.Name = "test.txt"
	writer.Write([]byte("gzip.Writer example\n"))

}
