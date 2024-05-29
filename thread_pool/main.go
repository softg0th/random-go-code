package main

import (
	"fmt"
	"sync"
)

type ThreadPool interface {
	Start(q *UnboundedBlockingQueue)
	Submit(q *UnboundedBlockingQueue, task func())
	Stop()
}

type TP struct {
	threads int
	stopped bool
}

type UnboundedBlockingQueue struct {
	mu    sync.Mutex
	queue []func()
}

func NewUnboundedBlockingQueue() *UnboundedBlockingQueue {
	return &UnboundedBlockingQueue{
		queue: make([]func(), 0),
	}
}

func (q *UnboundedBlockingQueue) Put(item func()) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, item)
}

func (q *UnboundedBlockingQueue) Get() func() {
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
				task := q.Get()
				if task == nil {
					break
				}
				task()
			}
		}()
	}
	wg.Wait()
}

func (tp *TP) Submit(q *UnboundedBlockingQueue, task func()) {
	q.Put(task)
}

func (tp *TP) Stop() {
	tp.stopped = true
}

func taskPrintHelloWorld() {
	fmt.Println("Hello world!")
}

func taskCalcAB() {
	a := 2
	b := 3
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func taskFizz() {
	fmt.Println("Fizz")
}

func taskBuzz() {
	fmt.Println("Buzz")
}

func main() {
	tq := NewUnboundedBlockingQueue()
	tp := NewThreadPool(4)

	tp.Submit(tq, taskPrintHelloWorld)
	tp.Submit(tq, taskFizz)
	tp.Submit(tq, taskCalcAB)
	tp.Submit(tq, taskBuzz)
	tp.Submit(tq, taskBuzz)
	tp.Submit(tq, taskFizz)
	tp.Submit(tq, taskPrintHelloWorld)
	tp.Start(tq)
	tp.Stop()
}
