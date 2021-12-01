package main

import "fmt"

func f2(f func()) {
	fmt.Println("this is f2")
	f()
}

func f3(x, y int) {
	fmt.Println("this is f3")
	fmt.Println(x + y)
}

//闭包传参
func f4(f func(x, y int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

func main() {

	f4(f3, 1, 2)()

	//闭包函数
	//f1 := func(x,y int) {
	//	fmt.Println(x+y)
	//}
	//f1(1,2)

}
