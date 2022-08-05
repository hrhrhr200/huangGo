package main

import "fmt"

func main() {
	var l1 person
	l1.age = 15
	l1.name = "huangrui"
	l1.sex = true
	fmt.Println(l1.name)

	var l2 = new(person)
	l2.age = 13
	l2.name = "huangrui2"
	l2.sex = true
	fmt.Println(l2.name, (*l2).name, &l2)

	var l3 = &person{"huangrui3", 17, false}
	fmt.Println(l3.name, (*l3).name)

	/*
		方法的接收器
	*/
	stu := &student{"huangrui", 2}
	stu.learn()
	var stu2 student
	fmt.Println(&stu2)
	stu2.learn()

	var new1 NewNetWork
	new1.screen = "screen1"
	new1.keyboard = "key11"
	new1.wifi = "wifi1"
	fmt.Println(new1.cul(1, 2))
	fmt.Println(new1.screen)

	var c cat
	var d dog
	c.eat()
	c.sleep()
	d.eat()
	d.sleep()

	var l LearningEx
	l.learn()
	l.hear()
	l.look()
}

type person struct {
	name string
	age  int64
	sex  bool
}

type student struct {
	name  string
	class int
}

func (stu student) learn() {
	fmt.Printf("%s同学正在学习", stu.name)
}

/*
  go 的继承
*/
type Computer struct {
	screen   string
	keyboard string
}

func (cpu Computer) cul(num1, num2 int) int {
	return num1 + num2
}

/*
需要注意的是，Go中可以嵌入多个结构体，但是多个结构体不能有相同的方法，
如果有参数和方法名完全相同的方法，在编译的时候就会报错。
所以Go不存在嵌入多个结构体后，被嵌入的几个结构体有相同的方法，
最后不知道选择执行哪个方法的情况，多个结构体方法相同时，直接编译就会报错。
*/
type NewNetWork struct {
	Computer
	wifi string
}

/*
	多态
*/
type animal interface {
	eat()
	sleep()
}

type cat struct {
}

type dog struct {
}

func (c cat) eat() {
	fmt.Println("小貓在吃飯")
}

func (c cat) sleep() {
	fmt.Println("小貓在睡覺")
}

func (d dog) eat() {
	fmt.Println("小狗在吃飯")
}

func (d dog) sleep() {
	fmt.Println("小狗在睡覺")
}

type HearLearning interface {
	hear()
}

type LookLearning interface {
	look()
}

type Learning interface {
	HearLearning
	LookLearning
	learn()
}

type LearningEx struct {
}

func (l LearningEx) learn() {
	fmt.Println("正在學習中！！")
}

func (l LearningEx) hear() {
	fmt.Println("用耳朵去學習!!")
}

func (l LearningEx) look() {
	fmt.Println("用眼睛去看去學習!!")
}
