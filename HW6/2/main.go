package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	for i := 0; i < 8; i++ {
		go func() {
			fmt.Println("World")
		}()
	}
	for i := 0; ; i++ {
		if i == 0 {
			runtime.Gosched()
		}
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello")
	}
}
