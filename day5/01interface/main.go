package main

import "fmt"

type animal interface {
	move()
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

func (ct cat) eat(food string) {
	fmt.Println("猫吃", food)
}

func main() {
	var anm animal

	cat := cat{
		feet: 11,
		name: "木星",
	}

	anm = cat
	anm.eat("老鼠")
	fmt.Println(anm)
	//{木星 11}

}
