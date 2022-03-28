package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//waitGroup

func f1(i int)  {
	wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {

	for i := 0; i< 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	wg.Wait()
	fmt.Println("'-----'")
}
