package main

import "fmt"

//定义一个叫的类型
type speaker interface {
	speak() //只要实现了speak方法的变量都是sperker类型  方法签名
}
type dog struct{}
type cat struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(s speaker) {
	s.speak()
}

func main() {
	var d1 dog
	var c1 cat

	da(d1)
	da(c1)

}
