package main

import "fmt"

type Num int

func main() {
	var num Num
	num.Increase(100)
	fmt.Println(num)
}

func (num *Num) Increase(a Num) {
	*num += a
}
