package main

import "fmt"

func main() {
	//切片
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"huang", "lei", "mi", "wei",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]

	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)

	//切片文法

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(t[0])

	//切片的默认行为
	bs := []int{2, 3, 5, 7, 11, 13}

	bs = bs[1:4]
	fmt.Println(bs)

	bs = bs[:2]
	fmt.Println(bs)

	bs = bs[1:]
	fmt.Println(bs)

	bs = bs[:3]
	fmt.Println(bs)

	//切片的长度和容量
	cs := []int{2, 3, 5, 7, 11, 13}
	printSlice(cs)

	//截取切片使其长度为0
	cs = cs[:0]
	printSlice(cs)

	//拓展其长度
	cs = cs[1:4]
	printSlice(cs)

	//舍弃前两个值
	cs = cs[2:]
	printSlice(cs)

	//nil切片
	var es []int
	fmt.Println(es, len(es), cap(es))
	if es == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
