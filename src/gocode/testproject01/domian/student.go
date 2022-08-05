package domian

type student struct {
	Age  int
	Name string
}

//工厂模式
func GetStudent(n int, name string) *student {
	return &student{n, name}
}
