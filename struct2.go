package main

import "fmt"

type Student struct {
	Id    int64
	Name  string
	Male  bool
	Score float64
}

func NewStudent(id int64, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

func (s Student) GetName() string {
	return s.Name
}

func main() {
	student := NewStudent(1, "joe", true, 64.1)
	fmt.Println(student)

	name := student.GetName()
	fmt.Println(name)
}
