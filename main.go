package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆
	s.Name = name
	s.Age = age
	return s
}
func main() {
	StudentRegister("Jim", 18)

	s := "Escape"
	fmt.Println(s)

	s1 := make([]int, 10000, 10000)
	for index, _ := range s {
		s1[index] = index
	}
}
