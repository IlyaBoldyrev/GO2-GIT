package main

import (
	"testing"
)

func BenchmarkWrite10Read90M(b *testing.B) {
	var m = newM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Write10Read90M()
			}
		})
	})
}

func BenchmarkWrite50Read50M(b *testing.B) {
	var m = newM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Write50Read50M()
			}
		})
	})
}

func BenchmarkWrite90Read10M(b *testing.B) {
	var m = newM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Write90Read10M()
			}
		})
	})
}
func BenchmarkWrite10Read90RW(b *testing.B) {
	var rw = newRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rw.Write10Read90RW()
			}
		})
	})
}
func BenchmarkWrite50Read50RW(b *testing.B) {
	var rw = newRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rw.Write50Read50RW()
			}
		})
	})
}
func BenchmarkWrite90Read10RW(b *testing.B) {
	var rw = newRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rw.Write90Read10RW()
			}
		})
	})
}
