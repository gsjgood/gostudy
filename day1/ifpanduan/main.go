package main

import "fmt"

func main()  {

	 age := 11

	 if age < 18 {
	 	fmt.Println("小学生")
	 } else if age == 18 {
		 fmt.Println("18岁的小学生")
	 } else {
		 fmt.Println("大于18的小学生")
	 }


	 for i := 0 ; i< 10;i++ {
	 	fmt.Println(i)

	 	if i == 4 {
	 		goto ssss  //相当于 break
		}
	 }
ssss: //指定标签
	fmt.Println("跳到指定标签")
	 s1 := "我是一个shel"

	 for _,value :=  range s1 {
		fmt.Printf("%c", value)
	 }


}