package main

import (
	"fmt"
	"gocode/testproject01/demo2/extm"
)

func main() {
	var s = extm.NewStudent("huangrui")
	fmt.Println(s.Name)
	s.SetAge(-1)
	s.SetAge(5)
	fmt.Println(s.GetAget())
}
