package main

import (
	"encoding/json"
	"fmt"
)

//go 语言中如果标识符首字母是大写的 就表示对外部可见 （共有的 暴露的）
type dog struct {
	name string
}

type dog1 struct {
	Name string `json:"name",db:"nAme"` //json 格式下重定义 db 下重定义
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//------------
//方法是作用域特定类型的函数
//接受者表示的是调用改方法的具体类型变量， 多用类型名 -> (d) 首字母小写表示
//直接收
func (d dog) wang() {
	fmt.Println("%s:汪汪汪：", d.name)
}

//指针接受   //如果一个方法是用来指针接受 那么为了保证一致 最好是 都变成指针接受  一般都是用 指针接受
func (d *dog) zhuazi() {
	d.name = d.name + "带上爪子"
	fmt.Println("%s:抓东西：", d.name)
}

//------------

//给自定义类型加方法
type myInt int

func (mi myInt) myIntHH() {
	fmt.Println("自定义类型方法")
}

//------------
//结构体匿名字段

type person struct {
	string
	int
}

//嵌套 结构体

type person1 struct {
	name    string
	age     int
	comy    company
	address //匿名嵌套
}

type address struct {
	province string
	city     string
}
type company struct {
	name     string
	province string
	city     string
}

//------------
//继承 结构体

type animal struct {
	name string
}

func (ani animal) move() {
	fmt.Println("会动")
}

type cat struct {
	name string
	animal
}

func (c cat) jiaosheng() {
	fmt.Println("喵喵喵")
}

func main() {
	d1 := newDog("旺财")
	d1.wang()
	d1.zhuazi()
	//%s:汪汪汪： 旺财
	//%s:抓东西： 旺财带上爪子
	//--------
	mm := myInt(100)
	mm.myIntHH()
	//自定义类型方法
	fmt.Println("------------")

	//匿名字段 不常用的  比较简单的场景
	p1 := person{
		"gsj",
		18,
	}

	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	//{gsj 18}
	//gsj
	//18
	fmt.Println("------------")

	//嵌套
	p2 := person1{
		name: "gsj",
		age:  18,
		address: address{
			province: "山西",
			city:     "运城",
		},
	}

	fmt.Println(p2)
	fmt.Println(p2.name, p2.address.city)
	fmt.Println(p2.city) //先找自己的 结构体里面的字段 找不到找匿名嵌套结构体中查找该字段 如果有多个相同嵌套的字段的话必须写上 嵌套的字段

	//{gsj 18 {  } {山西 运城}}
	//gsj 运城
	//运城

	fmt.Println("------------")
	//继承

	cat1 := cat{
		name: "木星",
		animal: animal{
			name: "猫",
		},
	}

	cat1.jiaosheng()
	cat1.move()
	//喵喵喵
	//会动
	fmt.Println("------------")
	fmt.Println("结构体与json")
	//把go语言的结构体 变量 -->json 格式的字符串
	//json 格式的字符串 -->  go语言的结构体 变量

	//序列化
	d2 := dog1{
		Name: "旺财",
	}
	b, err := json.Marshal(d2)
	if err != nil {
		fmt.Printf("错误%s\n", err)
		return
	}
	fmt.Println(string(b))
	//{"Name":"旺财"}

	//反序列化
	str := `{"name":"yff"}`
	var dog2 dog1
	json.Unmarshal([]byte(str), &dog2) //传指针 是为了能在json.Unmarshal 内部修改p2的值
	fmt.Printf("%#v\n", dog2)
	//main.dog1{Name:"yff"}

}
