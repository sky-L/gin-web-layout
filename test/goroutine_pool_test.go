package test

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"testing"
	"time"
)

func TestGoroutinePool(t *testing.T) {

	p, _ := ants.NewPool(10)

	_ = p.Submit(func() {
		fmt.Println("task")
		time.Sleep(time.Second)
	})

	fmt.Println(p.Free())

	p.Release()

	select {}
}
