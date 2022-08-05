package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["aaa"] = 42

	fmt.Println("aaa:", m["aaa"])

	m["aaa"] = 48

	fmt.Println("bbb:", m["aaa"])

	delete(m, "aaa")

	v, ok := m["aaa"]
	fmt.Println(v, ok)
}
