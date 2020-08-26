package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World.")
}

func Info(inter interface{}) {
	t := reflect.TypeOf(inter)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("类型错误")
		return
	}

	v:=reflect.ValueOf(inter)
	fmt.Println("Value:")

	for i := 0; i < t.NumField(); i++ {
		f:= t.Field(i)
		val:= v.Field(i).Interface()
		fmt.Println("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m:=t.Method(i)
		fmt.Println("%6s: %v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{1, "Joe", 13}
	Info(user)
}