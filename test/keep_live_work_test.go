package test

import (
	"fmt"
	"testing"
	"time"
)

type Worker struct {
	Id  int
	err error
}

func (wk *Worker) Work(failChannel chan<- *Worker) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("panic by %+v", r)
			}
		} else {
			wk.err = err
		}
		failChannel <- wk
	}()

	time.Sleep(1 * time.Second)
	panic("我退出了")
}

type WorkerManager struct {
	FailWork  chan *Worker
	WorkerNum int
}

func NewWorkerManager(workerNum int) *WorkerManager {
	return &WorkerManager{
		make(chan *Worker, workerNum),
		workerNum,
	}
}

func (w *WorkerManager) StartWorkerPool() {
	for i := 0; i < w.WorkerNum; i++ {
		wk := &Worker{i, nil}
		go wk.Work(w.FailWork)
	}
	w.keepLive()
}

func (w *WorkerManager) keepLive() {
	for wk := range w.FailWork {
		fmt.Println("死亡进程", wk.Id, wk.err.Error())
		wk.err = nil
		fmt.Println("复活进程", wk.Id)
		go wk.Work(w.FailWork)
	}
}

func TestKeepLiveWork(t *testing.T) {
	wp := NewWorkerManager(2)
	wp.StartWorkerPool()
}
