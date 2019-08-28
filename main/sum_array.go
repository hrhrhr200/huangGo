package main

import "fmt"

func sum2(a []float32) (sum float32) {
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return
}

func main() {
	/*a := []float32{2.0,3.0,4.0,5.0}
	sum := sum2(a)
	fmt.Println(sum)

	b := make([]int,5)

	min := minSlice(b)
	fmt.Println(min)*/

	s := []int{1, 2, 3, 4, 5, 6}
	//copyslice(s,5)

	s = RemoveStringSlice(0, 3, s)
	fmt.Println(s)
}

func minSlice(aaa []int) (min int) {
	min = aaa[0]
	for _, v := range aaa {
		if v < min {
			min = v
		}
	}
	return min
}

func copyslice(s []int, factor int) {
	fmt.Println(len(s))
	fmt.Println(s)
	news := make([]int, len(s)*factor)

	copy(news, s)

	s = news

	fmt.Println(len(s))
	fmt.Println(s)
}

func filter(s []int, fn func(int) bool) (result []int) {
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return
}

func even(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func RemoveStringSlice(start, end int, slice []int) []int {
	result := make([]int, len(slice)-(end-start)-1)
	at := copy(result, slice[:start])
	copy(result[at:], slice[end+1:])
	return result

}
