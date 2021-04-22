package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestId, err := uuid.NewRandom()
	if err != nil {
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "serve %s /\t%b\n", r.Method, requestId)
	w.Write([]byte("http.ResponseWriter sample"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
