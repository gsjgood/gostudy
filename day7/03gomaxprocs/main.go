package main

import (
	"fmt"
	"runtime"
	"sync"
)

//waitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A :%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B :%d\n", i)
	}
}

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(2) //默认 CPu的逻辑核心数 默认跑完整个cpu

	wg.Add(2)

	go a()
	go b()
	wg.Wait()

}
