package extm2

import "fmt"

type Animal struct {
	Age   int32
	Heavy float32
}

func (a *Animal) Hello() {
	fmt.Printf("%v 几岁 %v 多重\n", a.Age, a.Heavy)
}

type Cat struct {
	*Animal
}

type Dog struct {
	Animal
}
