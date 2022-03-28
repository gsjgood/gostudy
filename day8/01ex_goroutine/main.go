package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type job struct {
	value int64
}
type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func g1(jobChan chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		jobChan <- newJob
	}
}

func workerGoroutine(j <-chan *job, r chan<- *result) {
	//从jobchan中去除随机数计算各位数的和，讲解过发送到 r
	defer wg.Done()
	for {
		job := <-j
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		newResult := &result{
			job:    job,
			result: sum,
		}

		r <- newResult
	}

}

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	go g1(jobChan)
	wg.Add(24)
	//开启24个gone执行
	for i := 0; i < 24; i++ {
		go workerGoroutine(jobChan, resultChan)
	}

	//从resultchan去除结果并打印到终端输出
	for ret := range resultChan {
		fmt.Printf("value:%d sum:%d\n", ret.job.value, ret.result)
	}
	wg.Wait()

}
