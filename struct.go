package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	a := &person{
		Name: "joe",
		Age: 19,
	}

	a.Age = 199
	fmt.Println(a)
	A(a)
	fmt.Println(a)
}

func A(p *person) {
	p.Age = 20
	fmt.Println("A", p)
}
