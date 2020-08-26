package main

import "fmt"

func main() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
		//close(c)
	}()
	<-c

	//fmt.Println(len(c))
	//for v := range c {
	//	fmt.Println(v)
	//}
}
