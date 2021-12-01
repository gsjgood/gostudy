package main

import "fmt"

//type

type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {

	var m myInt
	m = 100
	fmt.Printf("%T\n", m)
	fmt.Println(m)

	//main.myInt
	//100
	var n yourInt
	n = 100
	fmt.Printf("%T\n", n)
	fmt.Println(n)
	//int
	//100

}
