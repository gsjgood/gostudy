package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {
	specialChar := []byte{' ', ',', ';', '!', '@', '$', '%', '@', '&', '*', '(', ')', '[', ']', '<', '>', '?', '.', '~', '`'}
	keyword := "宝马发动&&机$异响"
	for _, v := range specialChar {
		keyword = strings.ReplaceAll(keyword, string(v), "")
	}

	fmt.Println(keyword)

	var p1 = new(person)
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	//&{ 0}
	//*main.person
}
