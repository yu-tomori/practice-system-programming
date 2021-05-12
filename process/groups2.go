package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ユーザID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %d\n", groups)
}
