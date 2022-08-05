package main

import "fmt"

func printSlicea(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var i []int
	printSlicea(i)

	i = append(i, 0)
	printSlicea(i)

	i = append(i, 1)
	printSlicea(i)

	i = append(i, 2, 3, 4)
	printSlicea(i)
}
