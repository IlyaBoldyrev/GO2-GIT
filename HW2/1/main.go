//1
//
//2
//
//panicCall()
//
//3
package main

import (
	"fmt"
)

//Function panicCall generates panic implicitly
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
