package main

import "fmt"

func arrayv(a [3]int) {
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
}

func getarr() [15]int {
	var a = [15]int{}
	for i := 0; i < 15; i++ {
		a[i] = i
	}
	return a
}

func main() {
	/*a := [3]int{4,5,6}
	arrayv(a)
	fmt.Println(a)*/
	/*a := getarr()
	fmt.Println(a)*/
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
		fmt.Println(item)
	}
	fmt.Println(items)
}
