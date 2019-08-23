package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSliceA("a", a)

	b := make([]int, 0, 5)
	printSliceA("b", b)

	c := b[:2]
	printSliceA("c", c)

	d := c[2:5]
	printSliceA("d", d)
}

func printSliceA(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
