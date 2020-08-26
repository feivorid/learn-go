package main

import "fmt"

func main() {
	//b := 1
	//a := 1
	//fmt.Println(b + a)

	//var c int = 65
	//d := strconv.Itoa(c)
	//fmt.Println(d)
	//fmt.Println("Hello World!")

	a := [...]int{1,11,9,4,3}
	fmt.Println(a)

	l := len(a)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)
}
