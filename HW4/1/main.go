package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		j     int
		wg    sync.WaitGroup
		wp    = make(chan int, 8)
		mutex sync.Mutex
		arr   = make([]int, 0)
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
				mutex.Lock()
				j++
				mutex.Unlock()
			}()
		}
		wg.Wait()
		if j != 1000 { // если итоговое число не 1000, записываем номер прогона в слайс
			arr = append(arr, k)
		}
	}
	fmt.Println(arr) // выводим слайс, если он пуст, значит, мы всегда получали 1000
}
