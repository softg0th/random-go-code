package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Mutexx interface {
	Lock()
	Unlock()
	TestLock()
}

type Mutex struct {
	locked int32
}

func (m *Mutex) Lock() {
	currentMutexState := atomic.LoadInt32(&m.locked)

	if currentMutexState == 0 {
		atomic.StoreInt32(&m.locked, 1)
	} else {
		err := errors.New("can't lock locked mutex")
		fmt.Println(err)
		return
	}
}

func (m *Mutex) Unlock() {
	currentMutexState := atomic.LoadInt32(&m.locked)

	if currentMutexState == 1 {
		atomic.StoreInt32(&m.locked, 0)
	} else {
		err := errors.New("can't unlock unlocked mutex")
		fmt.Println(err)
		return
	}
}

func (m *Mutex) lockedOrNot() bool {
	if m.locked == 0 {
		return false
	} else {
		return true
	}
}

type testData struct {
	value int
	mu    *Mutex
}

func (t *testData) setData(newValue int) bool {
	if t.mu.lockedOrNot() {
		err := errors.New("can't store while mutex locked")
		fmt.Println(err)
		return false
	} else {
		t.mu.Lock()
		t.value = newValue
		fmt.Printf("%d successfully changed", t.value)
		t.mu.Unlock()
		return true
	}
}

func testCase(value int, mu *Mutex) {
	fmt.Printf("%d entered loop\n", value)

	for {
		td := &testData{mu: mu}
		fmt.Printf("%d current state of mutex\n", atomic.LoadInt32(&mu.locked))
		min := int64(1000000000)
		max := int64(10000000000)
		randomNumber := rand.Int63n(max-min+1) + min

		if td.setData(value) {
			fmt.Printf("%d was stored\n", value)
			time.Sleep(time.Duration(randomNumber))
			fmt.Printf("%d unlocked mutex\n", value)
		} else {
			time.Sleep(time.Duration(randomNumber))
			fmt.Printf("%d blocked mutex\n", td.value)
		}
	}
}

func main() {
	fmt.Println("start")
	mu := &Mutex{}
	for i := 5; i < 10; i++ {
		go testCase(i, mu)
	}
	select {}
}
