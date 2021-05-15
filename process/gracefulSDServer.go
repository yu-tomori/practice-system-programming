package main

import (
	"context"
	"fmt"
	"github.com/lestrrat/go-server-starter/listener"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	// listener.ListenAll:
	// 	listenAll parses environment variable SERVER_STARTER_PORT,
	// 	and creates net.Listener objects.
	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server pid: %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])

	<-signals
	server.Shutdown(context.Background())
}
