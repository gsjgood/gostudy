package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作
var x int64
var wg sync.WaitGroup

func add() {

	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {

	//wg.Add(1)
	//
	//for i := 0; i < 10000; i++ {
	//	go add()
	//}
	//
	//wg.Wait()

	//比较并交换
	ok := atomic.CompareAndSwapInt64(&x, 0, 200)
	fmt.Println(ok, x)
}
