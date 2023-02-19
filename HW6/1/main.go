package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		m  sync.Mutex
		ch = make(chan int, 5)
		j  = 0
	)
	for i := 0; i < 5; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			j++
			ch <- j
		}()
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
