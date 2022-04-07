package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)
import "golang.org/x/time/rate"

func TestRateLimit(t *testing.T) {

	limiter := rate.NewLimiter(rate.Limit(1), 1)

	for {
		err := limiter.Wait(context.Background())

		if err != nil {
			fmt.Println("超了")
			return
		}
		fmt.Println("gogo")
	}
}

func TestSelect(t *testing.T) {
	//rand.Seed(time.Now().Unix())

	/**
	2
	0
	2
	2
	1
	0
	*/

	for {
		fmt.Println(rand.Intn(3))
		time.Sleep(1 * time.Second)
	}

	//ch := make(chan int)
	//go func() {
	//	for {
	//		select {
	//		case ch <- 0:
	//		case ch <- 1:
	//		}
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//
	//for v := range ch {
	//	fmt.Println(v)
	//}

}
