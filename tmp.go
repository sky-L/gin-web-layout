package main

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

func (w *WaitGroupPool) Add() {
	w.ch <- struct{}{}
	w.wg.Add(1)
}

func (w *WaitGroupPool) Done() {
	<-w.ch
	w.wg.Done()
}

func (w *WaitGroupPool) Wait() {
	w.wg.Wait()
}
