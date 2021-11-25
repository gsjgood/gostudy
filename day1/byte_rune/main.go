package main

import (
	"fmt"
	"unicode"
)

func main()  {

	s := "hello world"

	for _,c := range s {
		fmt.Printf("%c", c)
	}
	
	//字符串修改
	s1 := "我是一个桌子"
	s2 := []rune(s1) // 字符串转换为 切片
	s2[0] = '你'
	fmt.Println(string(s2)) //切片转换为字符串

	//类型转换
	num1 := 11
	fl1 := float64(num1)
	fmt.Println(fl1)
	fmt.Printf("%T", fl1) //查看类型   %T

	s4 := "我是一个abc"
	var count int

	for _,c := range s4 {
		fmt.Printf("type : %T\n", c)
		if unicode.Is(unicode.Han, c) {
			fmt.Println("找到中文\n")
			count++
		}
	}

	fmt.Println(count)

}
