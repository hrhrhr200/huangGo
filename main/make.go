package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSliceA("a", a) // a 5 5 [0,0,0,0,0]

	b := make([]int, 0, 5)
	printSliceA("b", b) // b 0 5 []

	c := b[:2]
	printSliceA("c", c) // c 2 5 [0,0]

	d := c[2:5]
	printSliceA("d", d) // d 3 3 [0,0,0]
}

func printSliceA(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
