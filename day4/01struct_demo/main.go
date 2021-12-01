package main

import "fmt"

//结构体占用一块 连续的内存空间
type x struct {
	a int8
	b int8
	c int8
}

func main() {

	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}

	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))

	//0x1400001c092
	//0x1400001c093
	//0x1400001c094

}
