package main

import (
	"fmt"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
	"io/ioutil"
	"strings"
)

func main() {
	// observerを作成
	emitter := make(chan interface{})
	source := observable.Observable(emitter)

	// イベントを受け取るobserverを作成
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			line := item.(string)
			if strings.HasPrefix(line, "func ") {
				fmt.Println(line)
			}
		},
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	// observeableとobserverを接続(subscribe)
	sub := source.Subscribe(watcher)

	// observableに値を投入
	go func() {
		content, err := ioutil.ReadFile("future_promise.go")
		if err != nil {
			emitter <- err
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- line
			}
		}
		close(emitter)
	}()

	// 終了待ち
	<-sub
}
