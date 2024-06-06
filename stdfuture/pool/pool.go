package pool

import (
	"std/future"
	"sync"
)

type ThreadPool interface {
	Start(q *UnboundedBlockingQueue)
	Submit(q *UnboundedBlockingQueue, task func() future.Promise)
	Stop()
}

type TP struct {
	threads int
	stopped bool
}

type UnboundedBlockingQueue struct {
	mu    sync.Mutex
	queue []func() future.Promise
}

func NewUnboundedBlockingQueue() *UnboundedBlockingQueue {
	return &UnboundedBlockingQueue{
		queue: make([]func() future.Promise, 0),
	}
}

func (q *UnboundedBlockingQueue) Put(item func() future.Promise) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, item)
}

func (q *UnboundedBlockingQueue) Get() func() future.Promise {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		return nil
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func NewThreadPool(count int) *TP {
	return &TP{
		threads: count,
		stopped: false,
	}
}

func (tp *TP) Start(q *UnboundedBlockingQueue) {
	var wg sync.WaitGroup

	for i := 0; i < tp.threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				taskFunc := q.Get()
				if taskFunc == nil {
					break
				}
				promise := taskFunc()
				fut := promise.MakeFuture()
				result, exception := fut.Get()
				if exception != nil {
					promise.SetValue(result)
				} else {
					promise.SetException(exception)
				}
			}
		}()
	}
	wg.Wait()
}

func (tp *TP) Submit(q *UnboundedBlockingQueue, task func() future.Promise) {
	q.Put(task)
	tp.Start(q)
}

func (tp *TP) Stop() {
	tp.stopped = true
}
