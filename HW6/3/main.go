package main

import (
	"fmt"
)

func main() {

	var (
		ch = make(chan int, 5)
		j  = 0
	)
	for i := 0; i < 5; i++ {
		go func() {
			j++
			ch <- j
		}()
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
