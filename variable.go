package main

import "fmt"

func main() {
	var (
		a int
		b string
		c bool
	)

	d := 10

	var arr [10]int

	var struct1 struct {
		f int
	}

	var point *int

	var map1 map[string]int

	_, nickName := GetName()

	fmt.Println(a, b, c, d, arr, struct1, point, map1, nickName)
}

func GetName() (userName, nickName string) {
	return "testUserName", "xxxxx"
}
