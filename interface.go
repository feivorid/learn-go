package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

type PhoneConnect struct {
	name string
}

func (pc PhoneConnect) Name() string {
	return pc.name
}

func (pc PhoneConnect) Connect() {
	fmt.Println("Connected:", pc.name)
}

func Disconnected(usb interface{}) {
	if pc, ok := usb.(PhoneConnect); ok {
		fmt.Println("Disconnected:", pc.name)
		return
	}

	switch v := usb.(type) {
	case PhoneConnect:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown Device")
	}
}

func main() {
	a := PhoneConnect{"Apple"}
	b := 1
	a.Connect()
	Disconnected(a)
	Disconnected(b)
}
