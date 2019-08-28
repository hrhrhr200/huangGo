package main

import "fmt"

type test1 struct {
	a float32
	int
	string
}

func main() {
	test := test1{5.0, 5, "aaa"}
	fmt.Println(test)
}
