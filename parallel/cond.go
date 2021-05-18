package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			mutex.Lock()
			defer mutex.Unlock()

			cond.Wait()
			fmt.Println(name)
		}(name)
	}

	fmt.Println("よーい")
	time.Sleep(time.Second)
	fmt.Println("どん！")
	// 待っているgoroutineを一斉に引き起こす
	cond.Broadcast()
	time.Sleep(time.Second)
}
