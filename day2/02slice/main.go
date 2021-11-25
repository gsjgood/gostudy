package main

import "fmt"

//切片 slice
	//定义 初始化 长度跟容量 由数组->切片 切片->切片
func main() {
	//切片的定义
	var slice1 []int
	var slice2 []string

	fmt.Println(slice1, slice2) //[] []
	fmt.Println(slice1 == nil)  // true  空间没有存储
	fmt.Println(slice2 == nil)  // true

	//初始化

	slice1 = []int{1,2,3}
	slice2 = []string{"字", "符"}
	fmt.Println(slice1, slice2)
	fmt.Println(slice1 == nil)  // false 空间存储了
	fmt.Println(slice2 == nil)  // false

	//长度跟容量
	fmt.Printf("len:%d, cap:%d\n",len(slice1), cap(slice1)) //len:3, cap:3

	// 由数组->切片
	arr := [...]int{1,2,3,5,6,7}
	slice3 := arr[0:3]
	fmt.Println(slice3) //[1 2]
	//几种方式   arr[:2] arr[2:] arr[: ] 很好理解 略
	//切片数组的容量是指底层数组的容量  容量：从数组的切片第一个元素到最后的元素数量
	fmt.Printf("len:%d, cap:%d arr len:%d\n",len(slice3), cap(slice3), len(arr)) //len:3, cap:6 arr len:6

	//切片->切片
	//切片还可以再切
	slice31 := slice3[1:]
	fmt.Println(slice31) //[2 3]

	//*切片是引用类型 数组变动而变动 指向底层的地址



}
