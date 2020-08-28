package main

import "fmt"

func GetNum() int {
	return 100
}

const (
	a = iota
	b = iota
	c = iota
)

const a1, a2, a3 = iota * 2, iota * 2, iota * 2

const (
	d = iota * 2
	e = iota * 2
	f = iota * 2
)

func main() {
	fmt.Println(a, b, c, a1, a2, a3, d, e, f)
}
