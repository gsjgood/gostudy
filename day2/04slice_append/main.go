package main

import "fmt"

//append 为切片追加元素
	//append 复制
func main()  {
	s1 := []int{1,2,3}
	fmt.Printf("s1:%v len:%d cap:%d\n", s1, len(s1), cap(s1)) //s1:[1 2 3] len:3 cap:3

	s1 = append(s1, 4)
	// 追加元素 放不下之后 会重新定义一个新的数组来放
	//必须用原 来的 切片来放值
	// 长度增加的 规则
	//小于1024 翻倍  大于 1024 1/4
	//源码位置  /$GOOROOT/src/runtime/slice.go
	fmt.Printf("s1:%v len:%d cap:%d\n", s1, len(s1), cap(s1)) //s1:[1 2 3 4] len:4 cap:6

	s2 := []int{7,8,9}
	s1 = append(s1, s2...)  //拼接切片 用  ... 表示拆开
	fmt.Printf("s1:%v len:%d cap:%d\n", s1, len(s1), cap(s1)) //s1:[1 2 3 4 7 8 9] len:7 cap:12

	//复制 切片
	//copy

	//删除切片
	fmt.Println(s1[1:])
	s1 = append(s1[:1], s1[2:]...)


}