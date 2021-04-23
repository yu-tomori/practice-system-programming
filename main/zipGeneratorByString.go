package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("zip.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	writer, err := zipWriter.Create("file.txt")
	r := strings.NewReader("say hello")
	io.Copy(writer, r)
}
