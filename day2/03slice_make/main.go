package main

import "fmt"

//make 函数 创建切片
	//切片的本质
func main()  {
	s1 := make([]int, 5, 10)

	fmt.Printf(" s1= %v len= %d cap=%d\n", s1, len(s1), cap(s1))//s1= [0 0 0 0 0] len= 5 cap=10

	//切片的本质
	//相当于一个框 ， 连续的一块内存 只能展示 相同类型的值  真正的数据 都是保存再底层的数组里面的
	//引用类型（所以不能比较 但是可以跟nil比较 申请完内容之后!= nil ）
	s11 := make([]int, 0)
	var s2 []int
	s3 := []int{}
	fmt.Println(s11 == nil) // false
	fmt.Println(s2 == nil) // true
	fmt.Println(s3 == nil) // false


}


