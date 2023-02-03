package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		n  int
		wg = sync.WaitGroup{}
	)
	fmt.Println("Put the value:")
	fmt.Scanln(&n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
		}(i)
	}
	wg.Wait()
}
