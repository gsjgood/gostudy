package main

import "fmt"

func f1() int {
	return 10
}

//函数可以作为参数类型
func f2(x func() int) {
	ret := x()

	fmt.Println(ret)
}
func xx(int, int) int {
	return 111
}
func f3(x func() int) func(int, int) int {
	ret := func(xx int, yy int) int {
		return 111
	}
	return ret
}
func main() {
	a := f1
	fmt.Printf("type:%T\n", a)

	f2(a)
	f2(f1)

}
