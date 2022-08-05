package main

import (
	"fmt"
	"gocode/testproject01/domian"
)

func main() {
	// s = domian.student{Age: 15, Name: "huangrui"}
	var s = domian.GetStudent(15, "HUANGRUI")
	fmt.Println(*s)
}
