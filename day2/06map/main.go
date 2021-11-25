package main

import "fmt"

//map类型
	//
func main()  {

	var m1 map[string]int
	fmt.Println(m1) //map[]
	//初始化后使用
	m1 = make(map[string]int, 10) //估算好容量  尽量避免动态扩容
	m1["葛帅军"] = 1000
	m1["葛帅军111"] = 100011
	fmt.Println(m1) //map[葛帅军:1000]

	value,err := m1["葛帅军"]
	if !err {
		fmt.Println(" ")
	} else {
		fmt.Println(value)
 	}

 	//删除
 	delete(m1, "葛帅军")
	fmt.Println(m1)

}
