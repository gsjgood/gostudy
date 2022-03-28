package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("你好啊", i)
}

func main() {

	for i := 0; i < 10000; i++ {

		go hello(i)
		//go func(i int) {
		//	fmt.Println(i)
		//}(i) //匿名函数
	}
	fmt.Println("main")

	time.Sleep(time.Second)
	//main 函数结束了由mian函数启动的gorouine也结束了
}
