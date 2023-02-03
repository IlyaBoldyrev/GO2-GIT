package main

import (
	"sync"
)

type objectM struct {
	sync.Mutex
	m map[int]struct{}
}

type objectRW struct {
	sync.RWMutex
	m map[int]struct{}
}

func (o *objectM) ReadObject(i int) bool {
	o.Lock()
	defer o.Unlock()
	_, ok := o.m[i]
	return ok
}

func (o *objectRW) ReadObject(i int) bool {
	o.Lock()
	defer o.Unlock()
	_, ok := o.m[i]
	return ok
}

func (o *objectM) WriteObject(i int) {
	o.Lock()
	defer o.Unlock()
	o.m[i] = struct{}{}
}

func (o *objectRW) WriteObject(i int) {
	o.Lock()
	defer o.Unlock()
	o.m[i] = struct{}{}
}

func newM() *objectM {
	return &objectM{
		m: map[int]struct{}{},
	}
}
func newRW() *objectRW {
	return &objectRW{
		m: map[int]struct{}{},
	}
}

func (m *objectM) Write10Read90M() {
	for i := 0; i < 10; i++ {
		if i == 0 {
			m.WriteObject(i)
		} else {
			m.ReadObject(i)
		}
	}
}
func (m *objectM) Write50Read50M() {
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			m.WriteObject(i)
		} else {
			m.ReadObject(i)
		}
	}
}
func (m *objectM) Write90Read10M() {
	for i := 0; i < 10; i++ {
		if i == 0 {
			m.ReadObject(i)
		} else {
			m.WriteObject(i)
		}
	}
}
func (rw *objectRW) Write10Read90RW() {
	for i := 0; i < 10; i++ {
		if i == 0 {
			rw.WriteObject(i)
		} else {
			rw.ReadObject(i)
		}
	}
}
func (rw *objectRW) Write50Read50RW() {
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			rw.WriteObject(i)
		} else {
			rw.ReadObject(i)
		}
	}
}
func (rw *objectRW) Write90Read10RW() {
	for i := 0; i < 10; i++ {
		if i == 0 {
			rw.ReadObject(i)
		} else {
			rw.WriteObject(i)
		}
	}
}
