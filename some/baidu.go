package some

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReqBaidu() {
	// 创建一个 HTTP 客户端
	client := http.Client{}

	// 发送 GET 请求
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}
