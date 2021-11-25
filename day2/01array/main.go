package main

import "fmt"

//数组
//存放 元素的容器
//必须制定元素的类型跟长度
    //数组的初始化   数据的遍历
func main()  {
    var a1 [2]bool
    var a2 [3]bool

    fmt.Printf(" a1 : %T a2 : %T", a1, a2)

    //数组的初始化  不初始化 默认位 0 （布尔值：fasle 字符串："" 整形跟浮点型： 0）
    a1 = [2]bool{true,true}
    fmt.Println(a1)

    //根据初始值 自动判断长度
    a11 := [...]int{1,2,3,4,5,6,7,8}
    fmt.Println(len(a11)) //8
    //自动补气默认值 或者根据 索引赋值初始值
    a3 := [3]int{1,2}
    fmt.Println(a3) //[1 2 0]
    a31 := [3]int{0:1, 1:3, 2:2}
    fmt.Println(a31) //[1 3 2]

    //数据的遍历
    son := [2]string{"狗蛋", "狗剩"}

    for i:=0;i<len(son);i++ {
        fmt.Println(son[i])
    }
    for Rank,name := range son{
        fmt.Println(Rank,  name)
    }

    //多维数组
    son2 := [2][2]string{
        [2]string{"大狗蛋", "小狗蛋"},
        [2]string{"大狗剩", "小狗剩"},
    }

    for _,v := range son2 {
        fmt.Println(v)
        for _, v1 := range v{
            fmt.Println(v1)
        }
    }

    //值类型  支付  == !=
    // [n]*T 指针数组 *[n]T  数组指针





}
