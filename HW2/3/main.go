package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := [1000000]*os.File{}
	var err error
	for i, j := range arr {
		j, err = os.Create("/home/ilya/GO2-GIT/HW1/3/new/" + strconv.Itoa(i) + ".txt")
		if err != nil {
			fmt.Println(i)
			fmt.Println(j)
		}
		defer j.Close()
	}
}
