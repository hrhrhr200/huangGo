package main

import "fmt"

type huangi interface {
	M()
}

type huangt struct {
	s string
}

func (t huangt) M() {
	fmt.Println(t.s)
}

func main() {
	var i huangi = huangt{"huang"}
	i.M()
}
