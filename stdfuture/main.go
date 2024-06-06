package main

import (
	"fmt"
	"std/future"
	"std/pool"
	"time"
)

func taskPrintHelloWorld() func() future.Promise {
	return func() future.Promise {
		promise := future.NewPromise()
		go func() {
			promise.SetValue("Hello world!")
		}()
		return promise
	}
}

func taskCalcAB() func() future.Promise {
	return func() future.Promise {
		promise := future.NewPromise()
		go func() {
			a := 2
			b := 3
			c := a + b
			fmt.Printf("%d + %d = %d\n", a, b, c)
			promise.SetValue(nil)
		}()
		return promise
	}
}

func taskFizz() func() future.Promise {
	return func() future.Promise {
		promise := future.NewPromise()
		go func() {
			promise.SetValue("Fizz")
		}()
		return promise
	}
}

func taskBuzz() func() future.Promise {
	return func() future.Promise {
		promise := future.NewPromise()
		go func() {
			fmt.Println("Buzz")
			promise.SetValue(nil)
		}()
		return promise
	}
}

func main() {
	tq := pool.NewUnboundedBlockingQueue()
	tp := pool.NewThreadPool(4)

	tp.Start(tq)

	tp.Submit(tq, taskPrintHelloWorld())
	tp.Submit(tq, taskFizz())
	tp.Submit(tq, taskCalcAB())
	tp.Submit(tq, taskBuzz())
	tp.Submit(tq, taskBuzz())
	tp.Submit(tq, taskFizz())
	tp.Submit(tq, taskPrintHelloWorld())
	fmt.Println("start")
	time.Sleep(5 * time.Second)
	tp.Stop()
}
