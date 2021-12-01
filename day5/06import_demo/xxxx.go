package main

import (
	"fmt"

	calcs "studygo/day5/05calc" //calcs 别名
)

//包的路径从GOPATH/src后面的路径开始写起，路径分隔符用/
//想被别的包调用的标识符都要首字母大写
//单行导入和多行导入
//导入包的时候可以指定别名
//导入包不想使用包内部的标识符需要使用匿名导入
//每个包带入的时候会自动执行一个名为init()的函数 没有参数｜返回值

func main() {
	ret := calcs.Add(1, 2)
	fmt.Println(ret)
}
