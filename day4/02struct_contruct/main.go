package main

import "fmt"

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

//构造函数
//返回的是结构体韩式结构体指针
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
//可以定义多个   以 最好以new开头

//func newPerson(name string, age int)  person {
func newPerson(name string, age int) *person { //指针结构体

	return &person{ //指针结构体
		//return person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {

	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("gsj", 18)
	p2 := newPerson("葛帅军", 23)
	fmt.Println(p1, p2)

	//{gsj 18} {葛帅军 23}

	d1 := newDog("旺财")
	fmt.Println(d1)
	//{旺财}

}
