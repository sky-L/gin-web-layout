package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// interview1  执行 2 个函数， 2 秒后结束， 其中一个耗时 2 秒，一个耗时 3 秒
func f1(c chan struct{}) {
	time.Sleep(2 * time.Second)
	c <- struct{}{}
}

func f2(c chan struct{}) {
	time.Sleep(3 * time.Second)
	c <- struct{}{}
}

func TestInterview1(t *testing.T) {
	c1 := make(chan struct{})
	c2 := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		go f1(c1)
		select {
		case <-ctx.Done():
			fmt.Println("time out")
		case <-c1:
			fmt.Println("c1 down")
		}
		c1 <- struct{}{}
	}()

	go func() {
		go f2(c2)
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
		case <-c2:
			fmt.Println("f2 down")
		}
	}()

	select {}
}
