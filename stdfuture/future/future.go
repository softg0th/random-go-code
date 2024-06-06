package future

import (
	"fmt"
	"time"
)

type Future interface {
	Get() (interface{}, interface{})
}

type FutureData struct {
	promise *PromiseData
}

func (fd *FutureData) Get() (interface{}, interface{}) {

	if fd.promise.state == "PENDING" {
		fmt.Println("waiting")
		time.Sleep(5000)
	}

	if fd.promise.state == "COMPLETED" {
		fmt.Println("completed")
		return fd.promise.result, nil
	}
	fmt.Println("failed")
	return nil, fd.promise.exception
}
