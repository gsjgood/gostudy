package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
func main() {
	//readFileByIoutil1()
	//readFileBybufio2()
	readFile3()
}

//读取文件
func readFile3() {
	fileObj, err := os.Open("./main.go")

	if err != nil {
		fmt.Println("err:%v\n", err)
		return
	}

	defer fileObj.Close()

	var tmp [128]byte

	for {
		n, err := fileObj.Read(tmp[:])

		if err != nil {
			fmt.Println("err:%v\n", err)
			return
		}

		fmt.Printf("读取了%d个字符", n)
		fmt.Println(string(tmp[:n]))

		if n < 128 {
			return
		}
	}
}

//bufio 读取文件
func readFileBybufio2() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("读取文件错误 err:%v", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("读取文件错误 err:%v", err)
			return
		}

		fmt.Println(line)
	}
}

//ioutil 读取文件
func readFileByIoutil1() {
	ret, err := ioutil.ReadFile("./main.go")

	if err != nil {
		fmt.Println("读取文件错误 err:%v", err)
		return
	}

	fmt.Println(string(ret))
}
