package main

import (
	"fmt"
	"time"
)

func panicCall() (err error) {
	defer func() {
		er := recover()
		if er != nil {
			t := time.Now()
			err = fmt.Errorf("%s %v", t, er)
		}
	}()
	fmt.Scanln(nil)
	return nil
}

func main() {
	err := panicCall()
	if err != nil {
		fmt.Println("Panic occured: ", err)
	}
}
