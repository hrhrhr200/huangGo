package main

import (
	"container/list"
	"fmt"
)

func maptest() {
	map1 := map[string]int{
		"Monday":   1,
		"Tuesday":  2,
		"Wedonday": 3,
		"Thursday": 4,
		"Friday":   5,
		"Saturday": 6,
		"Sunday":   7,
	}
	_, ok := map1["Tuesday"]
	fmt.Println(ok)

	_, ok2 := map1["Hollyday"]
	fmt.Println(ok2)
}

func main() {
	//maptest()

	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(101)
	lst.PushBack(102)

	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
