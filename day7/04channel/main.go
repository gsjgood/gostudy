package main

import (
	"fmt"
	"sync"
)

//channel

var a []int
var b chan int //需要制定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	b = make(chan int) //不带缓冲区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中渠道了", x)
	}()

	b <- 10 //卡住了 fatal error: all goroutines are asleep - deadlock! 没有 goroutine 接收
	fmt.Println("10发送到通道b中了")
}

func bufChannel() {
	b = make(chan int, 2) //带缓冲区通道的初始化
	b <- 10
	fmt.Println("发送了一个10")
	b <- 20
	fmt.Println("发送了一个20")
	x := <-b
	fmt.Println("通道取到一个", x)
}

//单向通道
func sinChan(ch1 chan<- int) {
	ch1 <- 111
}

func main() {

	//bufChannel()
}
