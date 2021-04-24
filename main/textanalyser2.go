package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `一行目
二行目
三行目
`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
