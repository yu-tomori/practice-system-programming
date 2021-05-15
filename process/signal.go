package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	s := <-signals

	switch s {
	case syscall.SIGINT:
		fmt.Println("\ncatched SIGINT")
	case syscall.SIGTERM:
		fmt.Println("\ncatched SIGTERM")
	}
}
