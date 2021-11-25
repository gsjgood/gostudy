package main

import "fmt"

//指针操作

func main()  {

	//不存在指针操作   & *  互补操作
	//1 & 取地址
	str1 := "ss"
	p := &str1
	fmt.Println(p) //0x14000010240
	fmt.Printf("%T\n", p) //*string 字符串指针类型
	//2 * 根据地址取值

	str2 := *p
	fmt.Println(str2) //ss
	fmt.Printf("%T\n", str2) //string

	//new函数 申请指针地址
	var str3 = new(int)
	*str3 = 100
	fmt.Println(str3) //0x1400001c098
	fmt.Println(*str3) // 100

	//make 和new 区别
	// 都是用来申请内存的
	// new 很少用 一般都是用来给基本类型申请内存 string/int 返回对应类型的指针
	//make 给 slice 06map chan 申请内存的 返回对应类型本身




}
