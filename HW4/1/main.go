package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		j   int64
		wg  sync.WaitGroup
		wp  = make(chan int, 8)
		arr = make([]int, 0)
	)
	for k := 0; k < 1000; k++ { // тысяча проверок
		j = 0
		wg.Add(1000)
		for i := 0; i < 1000; i++ { // в этом цикле мы создаем тысячу горутин
			wp <- i
			go func() {
				defer func() {
					<-wp
					wg.Done()
				}()
				atomic.AddInt64(&j, 1)
			}()
		}
		wg.Wait()
		if j != 1000 { // если итоговое число не 1000, записываем номер прогона в слайс
			arr = append(arr, k)
		}
	}
	fmt.Println(arr) // выводим слайс, если он пуст, значит, мы всегда получали 1000
}
