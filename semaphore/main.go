package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Semaphore interface {
	Acquire(int)
	Release()
}

type countingSemaphore struct {
	counter  int64
	maxCount int64
	open     bool
}

func NewSemaphore(MaxCount int) *countingSemaphore {
	return &countingSemaphore{
		counter:  0,
		maxCount: int64(MaxCount),
		open:     true,
	}
}

func (c *countingSemaphore) Acquire(val int) {
	fmt.Printf("Current state: %d\n", c.counter)
	if c.counter >= c.maxCount {
		c.open = false

		for !c.open {
			var mutex sync.Mutex
			cond := sync.NewCond(&mutex)
			cond.L.Lock()
			cond.Wait()
		}
		fmt.Printf("%d approved\n", val)
		c.counter++
		return
	} else {
		fmt.Printf("%d approved\n", val)
		c.counter++
		return
	}
}

func (c *countingSemaphore) Release(val int) {
	fmt.Printf("%d leaved\n", val)
	if c.counter >= c.maxCount {
		c.open = true
		c.counter--
		return
	} else {
		c.counter--
		return
	}
}

func runTime(value int, cs *countingSemaphore) {
	fmt.Printf("%d entered loop\n", value)

	for {
		min := int64(1000000000)
		max := int64(10000000000)
		randomNumber := rand.Int63n(max-min+1) + min
		cs.Acquire(value)
		time.Sleep(time.Duration(randomNumber))
		cs.Release(value)
	}
}

func main() {
	fmt.Println("Start")
	sem := NewSemaphore(3)

	for i := 5; i < 10; i++ {
		go runTime(i, sem)
	}
	select {}
}
