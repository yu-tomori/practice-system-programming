package main

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	// map in memory
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer m.Unmap()

	// write bytes after fixing
	m[9] = 'X'
	m.Flush()

	// read
	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", fileData)
}
