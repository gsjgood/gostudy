package main

import (
	"fmt"
	"strings"
)

func main()  {

	str := "ssss"
	fmt.Println(str)
	str1 := `
		葛
		帅
		军
		`
	fmt.Println(str1)

	//字符串拼接
	str2 := "葛帅"
	str3 := "军"
	fmt.Println(str2+str3)
	ssdadsd := fmt.Sprintf("%s%s", str2, str3)
	fmt.Println(ssdadsd)
	path := "/Users/geshuaijun/go/src/studygo/day1"

	//字符串分割
	ret := strings.Split(path, "/")
	fmt.Println(ret)

	//判断前缀
	fmt.Println(strings.HasPrefix(path, "/Users"))
	fmt.Println(strings.HasPrefix(path, "/geshuaijun"))
	//判断后缀
	fmt.Println(strings.HasSuffix(path, "day1"))
	fmt.Println(strings.HasSuffix(path, "day2"))
	//判断 出现的位置
	fmt.Println(strings.Index(path, "e"))

	//拼接
	fmt.Println(strings.Join(ret, "*"))
}
