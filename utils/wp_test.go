package utils

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestWp(t *testing.T) {
	wgp := NewWaitGroupPool(10)

	go func() {
		for {
			fmt.Println(runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 1; i < 100; i++ {
		wgp.Add()
		go func() {
			defer wgp.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	wgp.Wait()
}
