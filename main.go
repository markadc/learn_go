package main

import (
	"demo/some"
	"fmt"
)

func main() {
	url := "https://www.baidu.com"
	text, err := some.ReqUrl(url)
	fmt.Println(len(text))
	if err != nil {
		fmt.Println("错误")
	}
}
