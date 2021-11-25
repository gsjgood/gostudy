package main

import "fmt"

//有返回值 可以 只写 return
func sum(x int, y int) (ret int) {
	//func sum(x int, y int) (ret int, res int) { //多个参数
	ret = x + y
	return
}

//没有返回值
func l1(x int, y int) {
	fmt.Println(x, y)
}

//有返回值 没有定义返回变量
func l2(x int, y int) int {
	fmt.Println(x, y)
	return x + y
}

//连续多个类型值相同可以省略
func l3(x, y int) int {
	fmt.Println(x, y)
	return x + y
}

//可变长参数 必须是最后一个
func l4(x string, y ...int) {
	fmt.Println(x, y)
}

func main() {

	sums := sum(1, 2)

	fmt.Println(sums)
}
