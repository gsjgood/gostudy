package main

import "fmt"

func main() {

	f1()
}

func f1() {
	fmt.Println("11111")

	//defer 用于函数结束之前释放资源 （文件具柄， 数据库函数 socket）
	//没错最后执行的 相同于 php 析构函数 的概念吧
	defer fmt.Println("555555")
	//defer fmt.Println("6666") //多个的话 顺序执行  针对函数

	fmt.Println("2222")
	fmt.Println("33333")
	fmt.Println("4444")

	//go  语言中的 return 不是原子操作
	//底层操作  1 返回值赋值 2，真正的RET 赋值
	// 有defer  1，返回值赋值 2，defer 3 ，真正的RET 赋值

	return

}
