package main

import (
	"fmt"
	"time"
)

func main() {
	//data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	data := []string{"a", "b"}
	// 带缓冲的chan 可以让producer不用开启协程
	// 如果无缓冲，则需要使用go producer来异步写入chan
	input := make(chan string, len(data))
	output := make(chan string)
	for i := 0; i < len(data); i++ {
		producer(data[i], input)
	}
	for i := 0; i < 10; i++ {
		go consumer(input, output)
	}
	// 使用output阻塞等待结果完成
	for i := 0; i < len(data); i++ {
		fmt.Println(i, <-output)
	}
	time.Sleep(1e9)
}

func producer(data string, input chan<- string) {
	fmt.Println("produce data", data)
	input <- data
}

func consumer(input <-chan string, output chan<- string) {
	for {
		msg := <-input
		if msg == "" {
			fmt.Println("nothinh")
			output <- "nothing"
			return
		}
		fmt.Println("consume data", msg)
		output <- msg + " done"
	}
}
