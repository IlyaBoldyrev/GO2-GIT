package main

import "sync"

func mutexUnlock(m *sync.Mutex) {
	m.Unlock()
}

func main() {
	var m = sync.Mutex{}
	go func() {
		m.Lock()
		defer mutexUnlock(&m)
	}()
}
