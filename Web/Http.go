package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	result, err := http.Get("https://baidu.com")

	if err != nil {
		fmt.Println("发起请求失败：", err)
		return
	}

	defer result.Body.Close()
	io.Copy(os.Stdout, result.Body)
}
