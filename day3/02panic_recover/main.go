package main

import "fmt"

func main() {

	var str1 string
	fmt.Scan(&str1)
	fmt.Println(str1)
	var (
		str2 string
		d1   int
	)
	fmt.Scanf("%s %d\n", &str2, &d1)
	fmt.Println(str2, d1)

	f1()
	f2()
	f3()
}

func f1() {
	fmt.Println("f1")
}
func f2() {

	defer func() {
		err := recover() //recover 必须在 defer 里面
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现错误！") //程序崩溃推出
	fmt.Println("f2")
}
func f3() {

	fmt.Println("f3")
}
