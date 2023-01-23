package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		sigChan     = make(chan os.Signal, 1)
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer func() {
		cancel()
		fmt.Printf("\nExit from main\n")
	}()
	signal.Notify(sigChan, syscall.SIGTERM)

	go func(ctx context.Context) {
		defer fmt.Printf("\nExit from goroutine")
		<-ctx.Done()
		return
	}(ctx)

	<-sigChan
	cancel()
	time.Sleep(50 * time.Millisecond)
	// Задержка, чтобы увидеть выход из горутины
}
