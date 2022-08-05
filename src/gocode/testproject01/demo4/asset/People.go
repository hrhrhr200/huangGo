package asset

import "fmt"

type Person interface {
	Sayhello()
}

type Chinese struct {
}

func (c Chinese) Sayhello() {
	fmt.Println("你好!")
}

func (c Chinese) Dance() {
	fmt.Println("跳养个")
}

func (a America) Play() {
	fmt.Println("play!!!")
}

type America struct {
}

func (a America) Sayhello() {
	fmt.Println("hello!")
}
