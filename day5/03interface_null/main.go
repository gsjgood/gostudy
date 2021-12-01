package main

import "fmt"

//空接口
//interface{}

//空接口作为函参数

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {

	var m1 map[string]interface{}

	m1 = make(map[string]interface{}, 16)
	m1["name"] = "gsj"
	m1["age"] = 18
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)
	//map[age:18 hobby:[唱 跳 rap] name:gsj]

	show(false)
	show(nil)
	show(m1)
	//type:bool value:false
	//type:<nil> value:<nil>
	//type:map[string]interface {} value:map[age:18 hobby:[唱 跳 rap] name:gsj]
}
