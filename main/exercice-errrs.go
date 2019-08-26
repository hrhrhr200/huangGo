package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v\n", float64(e))
}

func sqrt2(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		return 0, nil
	}
}

func main() {
	fmt.Println(sqrt2(2))
	fmt.Println(sqrt2(-2))
}
