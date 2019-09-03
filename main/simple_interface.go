package main

import "fmt"

type Simpler interface {
	get() int
	set(i int)
}

type simple struct {
	s int
}

func (s *simple) get() int {
	return s.s
}

func (s *simple) set(i int) {
	s.s = i
}

func test5(s1 Simpler) {
	s1.get()
	s1.set(5)
}

type rsimple struct {
	a int
}

func (s *rsimple) get() int {
	return s.a
}

func (s *rsimple) set(i int) {
	s.a = i
}

func main() {
	var s Simpler = &simple{2}
	fi(s)

	var a Simpler = &rsimple{3}
	fi(a)
}

func fi(simpler Simpler) {
	switch simpler.(type) {
	case *simple:
		fmt.Printf("Param #%T is a bbbb\n", simpler)
	case *rsimple:
		fmt.Printf("Param #%T is a aaaa\n", simpler)
	}
}
