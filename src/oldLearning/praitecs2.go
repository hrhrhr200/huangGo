package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Printf("Fields are: %q", strings.Fields("aaaaa   bbbbbb"))
	bm := WordCount("foo bar  baz")
	fmt.Println(bm)
}

func WordCount(s string) map[string]int {
	am := make(map[string]int)
	var strArray = strings.Fields(s)
	for i := 0; i < len(strArray); i++ {
		str := strArray[i]
		otherStrArray := strings.Split(str, "")
		for _, v := range otherStrArray {
			am[v] = am[v] + 1
		}
	}
	//return map[string]int{"x": 1}
	return am
}
