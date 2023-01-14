package main

import (
	"fmt"
)

func panicCall() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Panic happend:", err)
		}
	}()
	fmt.Scanln(nil)
}

func main() {
	panicCall()
}
