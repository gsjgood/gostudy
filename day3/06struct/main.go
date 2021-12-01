package main

import "fmt"

//结构体定义
//结构体是值类型

type person struct {
	name string
	age  int
}

func main() {
	var cla person
	cla.name = "gsj"
	cla.age = 18
	fmt.Println(cla)
	//{gsj 18}

	//匿名结构体
	var sscc struct {
		name string
		age  int
	}
	sscc.name = "hhh"
	sscc.age = 11
	fmt.Println(sscc)
	//{hhh 11}

	//指针类型的结构体
	cla2 := new(person)
	fmt.Println(cla2)
	fmt.Printf("%T %p\n", &cla2, cla2) //cla2 保存的值 就是一个内存地址
	fmt.Printf("%p\n", &cla2)
	//&{ 0}
	//**main.person  0x140000a0048
	//0x140000a8020

	//使用值的列表的初始化 值的顺序要和结构体定义时的 顺序 一致
	p4 := &person{
		"gsj",
		18,
	}
	fmt.Println(p4)

	//返回一个结构体变量的函数

}
