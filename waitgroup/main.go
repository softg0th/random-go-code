package main

import (
	"sync/atomic"
)

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

type WG struct {
	count int32
}

func NewWaitGroup() *WG {
	return &WG{
		count: 0,
	}
}

func (wg *WG) Add(delta int) {
	atomic.AddInt32(&wg.count, int32(delta))
}

func (wg *WG) Done() {
	if atomic.LoadInt32(&wg.count) > 0 {
		atomic.AddInt32(&wg.count, -1)
	} else {
		panic("negative counter")
	}
}

func (wg *WG) Wait() {
	for {
		if atomic.LoadInt32(&wg.count) == 0 {
			break
		}
	}
}

func main() {
	///
}
