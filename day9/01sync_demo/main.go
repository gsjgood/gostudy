package main

import (
	"fmt"
	"strconv"
	"sync"
)

//go内置的map不支持并发安全

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

var wg sync.WaitGroup

//func main() {
//
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(n int) {
//
//			key := strconv.Itoa(n)
//			set(key,n)
//			fmt.Printf("k:%v v:%v\n", key, get(key))
//			wg.Done()
//		}(i)
//
//		wg.Wait()
//	}
//}

var m2 = sync.Map{}

func main() {

	for i := 0; i < 21; i++ {

		wg.Add(1)

		go func(n int) {
			key := strconv.Itoa(n)

			m2.Store(key, n) //必须使用sync.map内置的strore方法设置键值对
			value, ok := m2.Load(key)
			if !ok {
				return
			}
			fmt.Printf("k:%v v:%v\n", key, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
