package utils

/**
* 控制并发数量
 */

import (
	"math"
	"sync"
)

type WaitGroupPool struct {
	ch chan struct{}
	wg *sync.WaitGroup
}

func NewWaitGroupPool(size int) *WaitGroupPool {
	if size == 0 {
		size = math.MaxInt32
	}
	return &WaitGroupPool{
		ch: make(chan struct{}, size),
		wg: &sync.WaitGroup{},
	}
}

func (p *WaitGroupPool) Add() {
	p.ch <- struct{}{}
	p.wg.Add(1)
}

func (p *WaitGroupPool) Done() {
	<-p.ch
	p.wg.Done()
}

func (p *WaitGroupPool) Wait() {
	p.wg.Wait()
}
