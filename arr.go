package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	for i, v := range arr {
		fmt.Println("Index:", i, "Value:", v)
	}

	for i := range arr {
		fmt.Println("Index:", i)
	}

	for _, v := range arr {
		fmt.Println("Value:", v)
	}
}
