package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	num := 0
	return func() int {
		num = num + 1
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
