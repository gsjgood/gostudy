package main

import "fmt"

//使用值接受者和指针接受者的区别
type animal interface {
	move()
	eater
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡跑")
}

func (ct cat) move() {
	fmt.Println("猫会跳")
}

func (ct *cat) eat(food string) {
	fmt.Println("猫吃", food)
}

//使用值接收者跟指针接收者实现接口的区别

//指针接受者实现接口只能存结构体指针类型的变量

//接口类型的关系

//同一个结构体可以实现多个接口
//接口也可以嵌套
func main() {

	var anm animal
	c1 := cat{"tom", 4}
	c2 := &cat{"黑猫警长", 4} //cat 类型的指针

	anm = &c1
	fmt.Println(anm) //{tom 4}
	anm = c2
	fmt.Println(anm) //&{黑猫警长 4}

}
