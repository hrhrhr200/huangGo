package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
}

func (c Car) numberOfWheels() int {
	return c.wheelCount
}

func (c Car) Start() {
	fmt.Println("gogogo start")
}

func (c Car) Stop() {
	fmt.Println("gogogo stop")
}

func (m Mercedes) sayHiToMerkel() {
	fmt.Println("hi Mercedes")
}

func main() {
	/*m := Mercedes{Car{nil,4}}
	fmt.Println(m.numberOfWheels())
	m.sayHiToMerkel()
	m.Start()
	m.Stop()*/
	e := Emploee{Person2{base{5}, "huang", "rui"}, 6}
	e.printId()
	e.setId(7)
	e.printId()
}

///////////练习2

type base struct {
	id int
}

func (b *base) Id() int {
	return b.id
}

func (b *base) setId(id int) {
	b.id = id
}

type Person2 struct {
	base
	FirstName string
	LastName  string
}

type Emploee struct {
	Person2
	salary int
}

func (e *Emploee) printId() {
	fmt.Println(e.id)
}
