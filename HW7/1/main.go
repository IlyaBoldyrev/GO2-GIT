package main

import (
	"fmt"
	"reflect"
)

type customStruct struct {
	key interface{}
}

func f(in *customStruct, values map[string]interface{}) error {
	p := reflect.ValueOf(&in.key)
	if !p.Elem().CanSet() {
		return fmt.Errorf("can`t set field of the map")
	}
	for i, v := range values {
		z := reflect.ValueOf(v)
		p.Elem().Set(z)
		fmt.Println(i, ": ", values, *in)
	}
	return nil
}

func main() {
	var (
		a map[string]interface{} = map[string]interface{}{
			"a": true,
			"b": 2,
			"c": 3.4,
			"d": "d",
			"e": struct{}{},
		}
		b customStruct
	)
	fmt.Println(a, b)
	fmt.Println()
	if err := f(&b, a); err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println(a, b)
}
