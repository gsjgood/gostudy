package main

import "fmt"



func main()  {
	//元素类型为map 的切片
	var s1  = make([]map[string]int, 10 , 100)

	//内部的map 需要做初始化
	s1[0] = make(map[string]int, 10)
	s1[0]["gsj"] = 1
	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["gsk"] = []int{1,2,3,5}

	fmt.Println(m1)
}
