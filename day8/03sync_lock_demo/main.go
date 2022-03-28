package main

import (
	"fmt"
	"sync"
)

//sync
var x int


func add()  {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x = x +1
		lock.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
