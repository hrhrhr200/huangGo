package main

import "fmt"

func splicStr(str string, index int) (s1, s2 string) {
	s1 = str[:index]
	s2 = str[index+1:]
	return
}

func main() {
	/*s1,s2 := splicStr("helloaaa",4)
	fmt.Println(s1,s2)*/
	/*str := convertstr("Google")
	fmt.Println(str)*/

	a := []int{4, 6, 9, 3, 20, 15, 7}
	//a = sortArr(a)
	a = mapFunc(mutinum, a)
	fmt.Println(a)
}

func convertstr(str string) string {
	byte2 := make([]byte, len(str))
	byte1 := []byte(str)
	for k, v := range byte1 {
		byte2[len(byte2)-1-k] = v
	}
	return string(byte2)
}

func sortArr(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				b := a[i]
				a[i] = a[j]
				a[j] = b
			}
		}
	}
	return a
}

func mapFunc(fn func(num int) int, numlist []int) []int {
	result := make([]int, len(numlist))
	for k, v := range numlist {
		v = fn(v)
		result[k] = v
	}
	return result
}

func mutinum(num int) int {
	return num * 10
}
