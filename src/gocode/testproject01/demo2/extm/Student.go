package extm

import "fmt"

type student struct {
	Name string
	age  int
}

func NewStudent(n string) *student {
	return &student{Name: n}
}

//定义age的set和get方法
func (s *student) SetAge(a int) {
	if a < 0 {
		fmt.Println("年龄不能小于0!")
		return
	}
	s.age = a
}

func (s *student) GetAget() int {
	return s.age
}
