package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//2021-12-01 10:48:05.70471 +0800 CST m=+0.000212251
	//2021
	//December
	//1
	//10
	//48
	//5
	//1638326885
	//1638326885704710000

	//时间戳
	ret := time.Unix(1638326885, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	//2021
	//December

	//时间间隔
	fmt.Println(now.Add(24 * time.Hour))
	fmt.Println(now.Add(-24 * time.Hour))
	//curr time 2021-12-01
	//2021-12-02 11:44:49.021394 +0800 CST m=+86400.000083709

	//定时器
	//timer := time.Tick(time.Second)
	//
	//for t := range timer {
	//	fmt.Println(t) //每秒执行一次
	//}

	//格式化时间
	fmt.Println(now.Format("2006-1-02"))
	//2021-12-01
	fmt.Println(now.Format(""))
	fmt.Println("-------------")
	fmt.Println("按照对应的格式解析字符串类型的时间")

	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-1-02", "2021-12-01")
	if err != nil {
		fmt.Println("err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	//2021-12-01 00:00:00 +0000 UTC
	//1638316800

	fmt.Println("----------------")
	fmt.Println("时区")
	timeObj2, err := time.Parse("2006-1-02", "2021-12-01")
	if err != nil {
		fmt.Println("err:%v\n", err)
		return
	}
	fmt.Println(timeObj2)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("err:%v\n", err)
		return
	}

	timeObj1, err := time.ParseInLocation("2006-1-02", "2021-12-01", loc)
	if err != nil {
		fmt.Println("err:%v\n", err)
		return
	}
	fmt.Println(timeObj1)
	//2021-12-01 00:00:00 +0000 UTC
	//2021-12-01 00:00:00 +0800 CST

	timeObj2, err2 := time.ParseInLocation("2006-01-02 15:04:05", time.Unix(time.Now().Unix(), 0).AddDate(0, 0, -30).Format("2006-01-02 ")+" 00:00:00", loc)
	if err2 != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	fmt.Println(timeObj2.Unix())

}
